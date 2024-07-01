package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Db *mongo.Database

func init() {
	opts := options.Client().ApplyURI("mongodb://152.42.209.24:27017")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	Db = client.Database("test")
	log.Print("Pinged your deployment. You successfully connected to MongoDB!")
}
