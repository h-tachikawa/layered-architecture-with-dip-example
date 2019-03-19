package handler

import (
	"layered-architecture-with-dip-example/domain/model"
	"layered-architecture-with-dip-example/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type VirtualLiverHandler interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
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
		return ectx.JSON(http.StatusInternalServerError, err)
	}
	return ectx.JSON(http.StatusOK, vl)
}

func (v *virtualLiverHandler) Create(ectx echo.Context) error {
	ctx := ectx.Request().Context()

	liver := new(model.VirtualLiver)
	err := ectx.Bind(liver)
	// TODO: validation

	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, err)
	}

	vl, err := v.u.Create(ctx, liver)

	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, err)
	}
	return ectx.JSON(http.StatusOK, vl)
}
