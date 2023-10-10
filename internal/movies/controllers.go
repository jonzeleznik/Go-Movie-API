package movies

import "github.com/gin-gonic/gin"

type MovieController struct {
	storage *MovieStorage
}

func NewMovieController(storage *MovieStorage) *MovieController {
	return &MovieController{
		storage: storage,
	}
}

func (t *MovieController) create(c *gin.Context) {
	// parse the request body
	var req MovieDB
	c.Bind(&req)

	// create the todo
	id, err := t.storage.createMovie(req)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err,
			"result": nil,
		})
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": id,
	})
}

func (t *MovieController) getAll(c *gin.Context) {
	// get all todos
	movies, err := t.storage.getAllTodos()
	if err != nil {
		c.JSON(400, gin.H{
			"error":  err,
			"result": nil,
		})
	}

	c.JSON(200, gin.H{
		"error":  nil,
		"result": movies,
	})
}
