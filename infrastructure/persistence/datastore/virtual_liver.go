package datastore

import (
	"context"
	"layered-architecture-with-dip-example/domain/model"
	"layered-architecture-with-dip-example/domain/repository"
	"os"

	w "go.mercari.io/datastore"
	"google.golang.org/appengine"

	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

const (
	virtualLiverKindName = "virtual-liver"
)

type virtualLiverRepository struct{}

func (v *virtualLiverRepository) Create(ctx context.Context, liver *model.VirtualLiver) (*model.VirtualLiver, error) {
	var client w.Client
	var err error

	if appengine.IsDevAppServer() {
		projectId := os.Getenv("DATASTORE_DATASET")
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
	} else {
		projectId := appengine.AppID(ctx)
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
	}

	if err != nil {
		panic(err)
	}

	key := client.IncompleteKey("virtual-liver", nil)
	_, err = client.Put(ctx, key, liver)

	return liver, nil
}

func NewVirtualLiverRepository() repository.VirtualLiverRepository {
	return &virtualLiverRepository{}
}

func (v *virtualLiverRepository) List(ctx context.Context) ([]*model.VirtualLiver, error) {
	livers := []*model.VirtualLiver{}

	var client w.Client
	var err error

	if appengine.IsDevAppServer() {
		projectId := os.Getenv("DATASTORE_DATASET")
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
	} else {
		projectId := appengine.AppID(ctx)
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
	}

	if err != nil {
		panic(err)
	}

	q := client.NewQuery(virtualLiverKindName)
	_, err = client.GetAll(ctx, q, &livers)

	return livers, nil
}
