package datastore

import (
	"context"
	"layered-architecture-with-dip-example/domain/model"
	"layered-architecture-with-dip-example/domain/repository"

	"github.com/google/uuid"
)

const (
	virtualLiverKindName = "VirtualLiver"
)

type virtualLiverRepository struct{}

func NewVirtualLiverRepository() repository.VirtualLiverRepository {
	return &virtualLiverRepository{}
}

func (v *virtualLiverRepository) Get(ctx context.Context, id string) (*model.VirtualLiver, error) {
	bm := NewBoomClient(ctx)

	liver := model.VirtualLiver{
		ID: id,
	}

	err := bm.Get(&liver)

	if err != nil {
		panic(err)
	}

	return &liver, nil
}

func (v *virtualLiverRepository) List(ctx context.Context) ([]*model.VirtualLiver, error) {
	bm := NewBoomClient(ctx)

	livers := []*model.VirtualLiver{}

	q := bm.NewQuery(virtualLiverKindName)
	_, err := bm.GetAll(q, &livers)

	if err != nil {
		panic(err)
	}

	return livers, nil
}

func (v *virtualLiverRepository) Create(ctx context.Context, liver *model.VirtualLiver) (*model.VirtualLiver, error) {
	bm := NewBoomClient(ctx)

	liver.ID = uuid.New().String()
	_, err := bm.Put(liver)

	if err != nil {
		panic(err)
	}

	return liver, nil
}
