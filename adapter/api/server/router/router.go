package router

import (
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	g := e.Group("/api/v1")
	//g.GET("/liver", func(c echo.Context) error {
	//})
}
