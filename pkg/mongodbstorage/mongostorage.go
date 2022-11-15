package mongodbstorage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage implements the Mongo storage backend
type Storage struct {
	// ClientURL is the mongodb  server address
	ClientURL string

	// Client is the MongoDB connection
	Client *mongo.Client

	//DataBase is the MongoDB database.
	DataBase string

	//Collection is the MongoDB Collection.
	Collection string
}

// Init initializes the mongodb storage
func (s *Storage) Init() error {
	if s.Client == nil {
		clientOptions := options.Client().ApplyURI(s.ClientURL)
		var errConnect error
		s.Client, errConnect = mongo.Connect(context.TODO(), clientOptions)
		if errConnect != nil {
			return errConnect
		}
	}

	if err := s.Client.Ping(context.Background(), nil); err != nil {
		return err
	}

	return nil
}
