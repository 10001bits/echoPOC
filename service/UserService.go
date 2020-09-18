package service

import (
	"context"
	"echoPoc/dao"
	"echoPoc/types"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	userdao dao.UserDaoInterface
}
var service *Service
func NewService() *Service  {
	service = &Service{dao.NewDao() }
	return service
}
func (service *Service)CreateUserService(u *types.User) (res *mongo.InsertOneResult,err error) {
	u.Name = "Mr."+u.Name
	if(u.Age<18){
		return nil, errors.New("Age less than 18")
	}
	result,err := service.userdao.AddUser(u)
	if err != nil {
		return result,err
	}
	return result,err
}
func (service *Service)UpdateUserService(u *types.User, id primitive.ObjectID) (user types.User,err error) {
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
	result := service.userdao.UpdateUser(filter,update,&returnOpt)
	if(result.Err()!=nil){
		return user,result.Err()
	}
	err = result.Decode(&user)
	return user,err
}
func (service *Service)GetUsersService() (Users[] types.User,err error){
	cur,errs := service.userdao.GetUsers(bson.M{})
	if errs != nil {
		return Users,errs
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var user types.User

		err := cur.Decode(&user)
		if err != nil {
			return Users,err
		}
		Users = append(Users, user)
	}
	return Users,err
}
func (service *Service)GetUserService(id primitive.ObjectID) (user types.User,err error) {
	filter := bson.M{"_id": id}
	result := service.userdao.GetUser(filter)
	if result.Err() != nil {
		return user,result.Err()
	}
	err = result.Decode(&user)
	return user,err
}
func (service *Service)DeleteUserService(id primitive.ObjectID)(res *mongo.DeleteResult,err error){
	filter := bson.M{"_id": id}
	 res, err = service.userdao.RemoveUser(filter)
	if err != nil {
		return res ,err
	}
	return res,err
}