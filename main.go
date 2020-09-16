package main

import (
	"context"
	"echoPoc/types"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)
var coll types.Coll

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://Ayush123:Ayush123@go-tinyurl-shard-00-00.hatxv.mongodb.net:27017,go-tinyurl-shard-00-01.hatxv.mongodb.net:27017,go-tinyurl-shard-00-02.hatxv.mongodb.net:27017/tinyUrl?ssl=true&replicaSet=atlas-4dlrbh-shard-0&authSource=admin&retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("go-tinyUrl").Collection("Users")
	coll =types.Coll{UserCollection: collection}
	e := echo.New()
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.PUT("/users/:id", UpdateUser)
	e.POST("/users",CreateUser)
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUser)
	e.DELETE("/users/:id", DeleteUser)

	e.Logger.Fatal(e.Start(":4000"))
}
func CreateUser(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest,u)
	}
	fmt.Println(u)
	result, err := coll.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return c.JSON(http.StatusBadRequest,u)
	}
	return c.JSON(http.StatusOK,result)
}
func UpdateUser(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		fmt.Println(u)
		return c.JSON(http.StatusBadRequest,u)
	}


	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": id}

	update := bson.D{
		{"$set",bson.D{
			{"name",u.Name},
			{"age",u.Age},
			{"email",u.Email},
		}},
	}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	err:= coll.UserCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt).Decode(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	return c.JSON(http.StatusOK,&u)
}
func DeleteUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": id}

	deleteResult, err := coll.UserCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	return c.JSON(http.StatusOK,deleteResult)
}
func GetUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": id}
	var user types.User
	err := coll.UserCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	return c.JSON(http.StatusOK,user)
}
func GetUsers(c echo.Context) error {
var users []types.User
	cur, err := coll.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.JSON(http.StatusBadRequest,users)
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var user types.User

		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK,users)
}