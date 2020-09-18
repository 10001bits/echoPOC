package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"time"
)

var (
	collection *mongo.Collection
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://Ayush123:Ayush123@go-tinyurl-shard-00-00.hatxv.mongodb.net:27017,go-tinyurl-shard-00-01.hatxv.mongodb.net:27017,go-tinyurl-shard-00-02.hatxv.mongodb.net:27017/tinyUrl?ssl=true&replicaSet=atlas-4dlrbh-shard-0&authSource=admin&retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDb")
	coll:= client.Database("go-tinyUrl").Collection("Users")
	 collection = coll
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys   : bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)
}
func GetCollection() *mongo.Collection {
	return collection
}
