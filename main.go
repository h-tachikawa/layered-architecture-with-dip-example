package main

import (
	"layered-architecture-with-dip-example/interfaces/api/server/router"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	router.Routes(e)
	http.Handle("/", e)
	appengine.Main()
	//e.Logger.Fatal(e.Start(":8080"))
}
