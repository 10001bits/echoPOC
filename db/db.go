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
	MongoDb *mongo.Database
)

func init() {
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
	MongoDb = client.Database("go-tinyUrl")
	coll:= MongoDb.Collection("Users")

	_, err = coll.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys   : bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)

}
func GetDb() (*mongo.Database){
	return MongoDb
}
func GetUserCollection()  *mongo.Collection {
	return MongoDb.Collection("Users")
}
