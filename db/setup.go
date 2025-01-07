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

func Setup() error {
	connString = config.Conf.MongoConnString
	dbLock = sync.Mutex{}

	fmt.Printf("Setting up mongo\n")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))

	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Mongo setup compleat\n")
	return nil
}
