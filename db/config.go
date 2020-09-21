package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

var (
	collection *mongo.Collection
)

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		os.Getenv("DB_URI"),
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
