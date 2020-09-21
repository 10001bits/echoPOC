package server

import (
	"net/http"
)

func Start()  {
	e:=Routes()
	http.Handle("/",e)
	e.Logger.Fatal(e.Start(":4000"))
}