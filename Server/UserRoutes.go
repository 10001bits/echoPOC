package Server

import (
	"echoPoc/Controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo){
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	 var Controllers = Controllers.NewController()
	e.PUT("/users/:id",Controllers.UpdateUser )
	e.POST("/users",Controllers.CreateUser)
	e.GET("/users", Controllers.GetUsers)
	e.GET("/users/:id", Controllers.GetUser)
	e.DELETE("/users/:id", Controllers.DeleteUser)
}
