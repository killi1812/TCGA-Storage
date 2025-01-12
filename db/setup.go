package db

import (
	"TCGA-storage/config"
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connString string
var dbLock sync.Mutex
var mongoDb *mongo.Client
var patients *mongo.Collection

const database = "tcga-storage"
const collection = "patients"

func Setup() error {
	connString = config.Conf.MongoConnString
	dbLock = sync.Mutex{}

	fmt.Printf("Setting up mongo\n")

	if err := connect(); err != nil {
		return err
	}

	fmt.Printf("Mongo setup compleat\n")
	return nil
}

func connect() error {
	dbLock.Lock()
	defer dbLock.Unlock()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connString))
	if err != nil {
		return err
	}
	mongoDb = client

	patients = client.Database(database).Collection(collection)

	return nil
}

func Disconnect() error {
	if err := mongoDb.Disconnect(context.Background()); err != nil {
		return err
	}
	return nil
}
