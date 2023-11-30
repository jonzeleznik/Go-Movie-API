package watchlist

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WatchListDB struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Title    string             `bson:"title" json:"title"`
	Movie_ID string             `bson:"movie_id" json:"movie_id"`
	TMDB_ID  int                `bson:"tmdb_id" json:"tmdb_id"`
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

func (s *WatchListStorage) getAllWatchList() ([]WatchListDB, error) {
	collection := s.db.Collection("watch_list")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	movies := make([]WatchListDB, 0)
	if err = cursor.All(context.TODO(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}
