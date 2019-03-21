package repository

import (
	"context"
	"layered-architecture-with-dip-example/domain/model"
)

type VirtualLiverRepository interface {
	Get(ctx context.Context, id string) (*model.VirtualLiver, error)
	List(ctx context.Context) ([]*model.VirtualLiver, error)
	Create(ctx context.Context, liver *model.VirtualLiver) (*model.VirtualLiver, error)
}
