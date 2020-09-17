package main

import (
	"echoPoc/Routes"
	"github.com/labstack/echo"
	"net/http"
)


func main() {
	e := echo.New()
	Routes.Routes(e)
	http.Handle("/",e)
	e.Logger.Fatal(e.Start(":4000"))
}
