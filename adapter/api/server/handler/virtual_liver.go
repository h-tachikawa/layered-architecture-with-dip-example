package handler

import (
	"layered-architecture-with-dip-example/usecase"
	"net/http"

	"google.golang.org/appengine"

	"github.com/labstack/echo"
)

type VirtualLiverHandler interface {
	List(ctx echo.Context) error
}

type virtualLiverHandler struct {
	u usecase.VirtualLiverUseCase
}

func NewVirtualLiverHandler(u usecase.VirtualLiverUseCase) VirtualLiverHandler {
	return &virtualLiverHandler{u}
}

func (v *virtualLiverHandler) List(ectx echo.Context) error {
	ctx := appengine.NewContext(ectx.Request())

	vl, err := v.u.List(ctx)

	if err != nil {
		// TODO: エラーハンドリング
	}
	return ectx.JSON(http.StatusOK, vl)
}
