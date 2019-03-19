package router

import (
	"layered-architecture-with-dip-example/registry"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	registry.InitRegistry()
	reg := registry.GetRegistry()
	g := e.Group("/api/v1")

	g.GET("/liver", func(ctx echo.Context) error {
		return reg.VirtualLiverHandler.List(ctx)
	})

	g.POST("/liver", func(ctx echo.Context) error {
		return reg.VirtualLiverHandler.Create(ctx)
	})
}
