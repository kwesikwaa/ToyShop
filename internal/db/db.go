package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

type MongoCollections string

const (
	ToysCollection MongoCollections = "toys"
)

const (
	url          = "mongodb://localhost:27017"
	DatabaseName = "toys-db"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(url)

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}
		clientInstance = client
		clientInstanceError = err
	})
	return clientInstance, clientInstanceError

}

func DisconnectMongo() {
	if err := clientInstance.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

}
