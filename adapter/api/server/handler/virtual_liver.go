package handler

import (
	"layered-architecture-with-dip-example/usecase"
	"net/http"

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
	ctx := ectx.Request().Context()

	vl, err := v.u.List(ctx)

	if err != nil {
		// TODO: エラーハンドリング
	}
	return ectx.JSON(http.StatusOK, vl)
}
