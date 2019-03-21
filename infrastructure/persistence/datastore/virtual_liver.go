package datastore

import (
	"context"
	"layered-architecture-with-dip-example/domain/model"
	"layered-architecture-with-dip-example/domain/repository"
	"os"

	"github.com/google/uuid"

	"go.mercari.io/datastore/boom"

	w "go.mercari.io/datastore"
	"google.golang.org/appengine"

	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

const (
	virtualLiverKindName = "VirtualLiver"
)

type virtualLiverRepository struct{}

func NewVirtualLiverRepository() repository.VirtualLiverRepository {
	return &virtualLiverRepository{}
}

func (v *virtualLiverRepository) Get(ctx context.Context, id string) (*model.VirtualLiver, error) {
	liver := model.VirtualLiver{
		ID: id,
	}

	var client w.Client
	var err error
	var bmcli *boom.Boom

	if appengine.IsDevAppServer() {
		projectId := os.Getenv("DATASTORE_DATASET")
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
		bmcli = boom.FromClient(ctx, client)
	} else {
		projectId := appengine.AppID(ctx)
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
		bmcli = boom.FromClient(ctx, client)
	}

	if err != nil {
		panic(err)
	}

	err = bmcli.Get(&liver)

	return &liver, nil
}

func (v *virtualLiverRepository) List(ctx context.Context) ([]*model.VirtualLiver, error) {
	livers := []*model.VirtualLiver{}

	var client w.Client
	var bmcli *boom.Boom
	var err error

	if appengine.IsDevAppServer() {
		projectId := os.Getenv("DATASTORE_DATASET")
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)

		if err != nil {
			panic(err)
		}

		bmcli = boom.FromClient(ctx, client)
	} else {
		projectId := appengine.AppID(ctx)
		opts := datastore.WithProjectID(projectId)
		client, err := clouddatastore.FromContext(ctx, opts)

		if err != nil {
			panic(err)
		}

		bmcli = boom.FromClient(ctx, client)
	}

	q := bmcli.NewQuery(virtualLiverKindName)
	_, err = bmcli.GetAll(q, &livers)

	if err != nil {
		panic(err)
	}

	return livers, nil
}

func (v *virtualLiverRepository) Create(ctx context.Context, liver *model.VirtualLiver) (*model.VirtualLiver, error) {
	var client w.Client
	var err error
	var bmcli *boom.Boom

	if appengine.IsDevAppServer() {
		projectId := os.Getenv("DATASTORE_DATASET")
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
		bmcli = boom.FromClient(ctx, client)
	} else {
		projectId := appengine.AppID(ctx)
		opts := datastore.WithProjectID(projectId)
		client, err = clouddatastore.FromContext(ctx, opts)
		bmcli = boom.FromClient(ctx, client)
	}

	if err != nil {
		panic(err)
	}

	liver.ID = uuid.New().String()
	_, err = bmcli.Put(liver)

	if err != nil {
		panic(err)
	}

	return liver, nil
}
