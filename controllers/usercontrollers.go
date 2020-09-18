package controllers

import (
	"echoPoc/beans"
	"echoPoc/service"
	"echoPoc/types"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)
type UserController struct {
	UserService service.UserService
}
var controller *UserController
func NewController() *UserController{
	controller = &UserController{service.NewService()}
	return controller
}
func (controller *UserController)CreateUser(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest,u)
	}
	result, err := controller.UserService.CreateUserService(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest,u)
	}
	return c.JSON(http.StatusOK,result)
}
func (controller *UserController)UpdateUser(c echo.Context) error {
	u := new(types.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest,u)
	}
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	user, err := controller.UserService.UpdateUserService(u, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	Response := beans.List1user(&user)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK,Response)
}
func (controller *UserController)DeleteUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	deleteResult, err := controller.UserService.DeleteUserService(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	return c.JSON(http.StatusOK,deleteResult)
}
func (controller *UserController)GetUser(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	user,err := controller.UserService.GetUserService(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest,id)
	}
	Response := beans.List1user(&user)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK,Response)
}
func (controller *UserController)GetUsers(c echo.Context) error {
	var users types.Users
	users, err := controller.UserService.GetUsersService()
	if err != nil {
		return c.JSON(http.StatusBadRequest,users)
	}
	Response := beans.Listalluser(&users)
	Response.Httpstatus = true
	return c.JSON(http.StatusOK,Response)
}