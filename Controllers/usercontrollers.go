package Controllers

import (
	"echoPoc/types"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func CreateUser(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest,u)
	}
	fmt.Println(u)
	result, err := u.AddUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest,u)
	}
	return c.JSON(http.StatusOK,result)
}
func UpdateUser(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
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

	err:= u.UpdateUser(filter,update,&returnOpt).Decode(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	return c.JSON(http.StatusOK,&u)
}
func DeleteUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": id}
	deleteResult, err := types.RemoveUser(filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	return c.JSON(http.StatusOK,deleteResult)
}
func GetUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": id}
	user,err := types.GetUser(filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	return c.JSON(http.StatusOK,user)
}
func GetUsers(c echo.Context) error {
	var users []types.User
	users, err := types.GetUsers(bson.M{})
	if err != nil {
		return c.JSON(http.StatusBadRequest,users)
	}
	return c.JSON(http.StatusOK,users)
}