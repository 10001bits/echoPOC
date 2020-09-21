package dao

import (
	"context"
	"echoPoc/db"
	"echoPoc/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type coll struct {
	UserCollection *mongo.Collection
}
var Collection *coll
func  NewDao() *coll{
	Collection = &coll{UserCollection: db.GetUserCollection()}
	return Collection
}

func (Collection *coll) AddUser(u *types.User) (result *mongo.InsertOneResult ,err error){
	result, err = Collection.UserCollection.InsertOne(context.TODO(), u)
	return result,err
}
func (Collection *coll)UpdateUser(filter interface{},update interface{},options *options.FindOneAndUpdateOptions) (result *mongo.SingleResult){
	result =Collection.UserCollection.FindOneAndUpdate(context.TODO(),filter,update,options)
	return result
}
func (Collection *coll)RemoveUser(filter interface{}) (result *mongo.DeleteResult,err error){
	result,err = Collection.UserCollection.DeleteOne(context.TODO(),filter)
	return result,err
}
func (Collection *coll)GetUser(filter interface{}) (result *mongo.SingleResult){
	result = Collection.UserCollection.FindOne(context.TODO(),filter)
	return result
}
func (Collection *coll)GetUsers(filter interface{}) (cur *mongo.Cursor ,err error){
	cur,err = Collection.UserCollection.Find(context.TODO(),filter)
	return cur,err
}
