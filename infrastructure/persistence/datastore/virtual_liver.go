package datastore

import (
	"context"
	"layered-architecture-with-dip-example/domain/model"
	"layered-architecture-with-dip-example/domain/repository"
	"os"

	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

const (
	virtualLiverKindName = "virtual-liver"
)

type virtualLiverRepository struct{}

func NewVirtualLiverRepository() repository.VirtualLiverRepository {
	return &virtualLiverRepository{}
}

func (v *virtualLiverRepository) List(ctx context.Context) ([]*model.VirtualLiver, error) {
	livers := []*model.VirtualLiver{}

	projectId := os.Getenv("DATASTORE_DATASET")
	opts := datastore.WithProjectID(projectId)
	client, err := clouddatastore.FromContext(ctx, opts)

	if err != nil {
		panic(err)
	}

	q := client.NewQuery(virtualLiverKindName)
	_, err = client.GetAll(ctx, q, &livers)

	return livers, nil
}
