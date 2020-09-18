package main

import (
	"echoPoc/server"
	"github.com/labstack/echo/v4"
	"net/http"
)


func main() {
	e := echo.New()
	server.Routes(e)
	http.Handle("/",e)
	e.Logger.Fatal(e.Start(":4000"))
}
