package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/jonzeleznik/Go-Movie-API/internal/movies"
	"github.com/jonzeleznik/Go-Movie-API/internal/storage"
	"github.com/jonzeleznik/Go-Movie-API/pkg/shutdown"
)

type EnvVars struct {
	MONGODB_URI  string
	MONGODB_NAME string
	PORT         string
}

func main() {
	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	var env EnvVars

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	env.MONGODB_URI = os.Getenv("MONGODB_URI")
	env.MONGODB_NAME = os.Getenv("MONGODB_NAME")
	env.PORT = os.Getenv("PORT")

	// run the server
	cleanup, err := run(env)

	// run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}

func run(env EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// start the server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// return a function to close the server and database
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env EnvVars) (*fiber.App, func(), error) {
	// init the storage
	db, err := storage.BootstrapMongo(env.MONGODB_URI, env.MONGODB_NAME, 10*time.Second)
	if err != nil {
		return nil, nil, err
	}

	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// create the user domain
	movieStore := movies.NewMovieStorage(db)
	movieController := movies.NewMovieController(movieStore)
	movies.AddMovieRoutes(app, movieController)

	return app, func() {
		storage.CloseMongo(db)
	}, nil
}
