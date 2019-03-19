package usecase

import (
	"context"
	"layered-architecture-with-dip-example/domain/model"
	"layered-architecture-with-dip-example/domain/repository"
)

type VirtualLiverUseCase interface {
	List(context.Context) ([]*model.VirtualLiver, error)
	Create(context.Context, *model.VirtualLiver) (*model.VirtualLiver, error)
}

type virtualLiverUseCase struct {
	repository.VirtualLiverRepository
}

func NewVirtualLiverUseCase(r repository.VirtualLiverRepository) VirtualLiverUseCase {
	return &virtualLiverUseCase{
		r,
	}
}

func (u *virtualLiverUseCase) List(ctx context.Context) ([]*model.VirtualLiver, error) {
	return u.VirtualLiverRepository.List(ctx)
}

func (u *virtualLiverUseCase) Create(ctx context.Context, liver *model.VirtualLiver) (*model.VirtualLiver, error) {
	return u.VirtualLiverRepository.Create(ctx, liver)
}
