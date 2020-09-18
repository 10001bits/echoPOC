package DAO

import (
	"echoPoc/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDaoInterface interface {
	AddUser(u *types.User) (result *mongo.InsertOneResult ,err error)
	UpdateUser(filter interface{},update interface{},options *options.FindOneAndUpdateOptions) (result *mongo.SingleResult)
	RemoveUser(filter interface{}) (result *mongo.DeleteResult,err error)
	GetUser(filter interface{}) (result *mongo.SingleResult)
	GetUsers(filter interface{}) (cur *mongo.Cursor ,err error)
}