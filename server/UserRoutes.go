package server

import (
	"echoPoc/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(e *echo.Echo){
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	 var Controllers = controllers.NewController()

	g:= e.Group("/api/v1/users",UserCtx)
	g.PUT("/:id",Controllers.UpdateUser )
	g.POST("",Controllers.CreateUser)
	g.GET("", Controllers.GetUsers)
	g.GET("/:id", Controllers.GetUser)
	g.DELETE("/:id", Controllers.DeleteUser)
}
func UserCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		return next(c)
	}
}