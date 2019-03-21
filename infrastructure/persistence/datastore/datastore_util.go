package datastore

import (
	"context"
	"os"

	"go.mercari.io/datastore"
	w "go.mercari.io/datastore"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore/clouddatastore"
	"google.golang.org/appengine"
)

func NewBoomClient(ctx context.Context) *boom.Boom {
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
	return bmcli
}
