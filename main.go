package main

import (
	"layered-architecture-with-dip-example/adapter/api/server/router"
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

func main() {
	e := echo.New()
	router.Routes(e)
	http.Handle("/", e)
	appengine.Main()
}
