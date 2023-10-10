package watchlist

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WatchListDB struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Title   string             `bson:"title" json:"title"`
	MovieID string             `bson:"movie_id" json:"movie_id"`
	TMDB_ID int                `bson:"tmdb_id" json:"tmdb_id"`
}

type WatchListStorage struct {
	db *mongo.Database
}

func NewWatchListStorage(db *mongo.Database) *WatchListStorage {
	return &WatchListStorage{
		db: db,
	}
}

func (s *WatchListStorage) createWatchList(movie WatchListDB) (string, error) {
	collection := s.db.Collection("watch_list")

	movie.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return "", err
	}

	// convert the object id to a string
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
