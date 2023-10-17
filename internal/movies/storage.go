package movies

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Genre struct {
	Id   int
	Name string
}

type MovieDB struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Title         string             `bson:"title" json:"title"`
	TMDB_ID       int                `bson:"tmdb_id" json:"tmdb_id"`
	IMDB_ID       string             `bson:"imdb_id" json:"imdb_id"`
	Overview      string             `bson:"overview" json:"overview"`
	Genre         []Genre            `bson:"genre" json:"genre"`
	Release_date  string             `bson:"release_date" json:"release_date"`
	Runtime       int                `bson:"Runtime" json:"Runtime"`
	Poster_path   string             `bson:"poster_path" json:"poster_path"`
	Backdrop_path string             `bson:"backdrop_path" json:"backdrop_path"`
}

type MovieStorage struct {
	db *mongo.Database
}

func NewMovieStorage(db *mongo.Database) *MovieStorage {
	return &MovieStorage{
		db: db,
	}
}

func (s *MovieStorage) createMovie(movie MovieDB) (string, error) {
	collection := s.db.Collection("movies")

	movie.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return "", err
	}

	// convert the object id to a string
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *MovieStorage) getAllMovies() ([]MovieDB, error) {
	collection := s.db.Collection("movies")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	movies := make([]MovieDB, 0)
	if err = cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *MovieStorage) searchMovies(title string) ([]MovieDB, error) {
	collection := s.db.Collection("movies")

	regex := `.*` + title + `.*`
	cursor, err := collection.Find(context.TODO(), bson.M{"title": bson.M{"$regex": regex}})
	if err != nil {
		return nil, err
	}

	movies := make([]MovieDB, 0)
	if err = cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}
