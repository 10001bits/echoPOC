package types
import (
	"context"
	"echoPoc/config"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name      string `json:"name" bson:"name"`
	Email	string	`json:"email" bson:"email"`
	Age 	int 	`json:"age" bson:"age"`
}

var UserCollection *mongo.Collection
func init(){
	config.Connect()
	UserCollection=config.GetCollection()
}
func (u *User)AddUser() (result *mongo.InsertOneResult ,err error){
	result, err = UserCollection.InsertOne(context.TODO(), u)
	return result,err
}
func (u *User)UpdateUser(filter interface{},update interface{},options *options.FindOneAndUpdateOptions) (result *mongo.SingleResult){
	result =UserCollection.FindOneAndUpdate(context.TODO(),filter,update,options)
	return result
}
func RemoveUser(filter interface{}) (result *mongo.DeleteResult,err error){
	result,err = UserCollection.DeleteOne(context.TODO(),filter)
	return result,err
}
func GetUser(filter interface{}) (user User,err error){
	err = UserCollection.FindOne(context.TODO(),filter).Decode(&user)
	return user,err
}
func GetUsers(filter interface{}) (users []User,err error){
	cur,err := UserCollection.Find(context.TODO(),filter)
	if err != nil {
		return users,err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var user User

		err := cur.Decode(&user)
		if err != nil {
			return users,err
		}
		users = append(users, user)
	}
	return users,err
}
