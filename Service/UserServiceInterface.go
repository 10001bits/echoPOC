package Service

import (
	"echoPoc/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface{
	CreateUserService(u *types.User) (res *mongo.InsertOneResult,err error)
	UpdateUserService(u *types.User, id primitive.ObjectID) (user types.User,err error)
	GetUsersService() (Users[] types.User,err error)
	GetUserService(id primitive.ObjectID) (user types.User,err error)
	DeleteUserService(id primitive.ObjectID)(res *mongo.DeleteResult,err error)
}
