package router

import (
	"layered-architecture-with-dip-example/domain/model"
	"layered-architecture-with-dip-example/registry"
	"net/http"
	"os"

	w "go.mercari.io/datastore"
	"google.golang.org/appengine"

	"go.mercari.io/datastore"

	"go.mercari.io/datastore/clouddatastore"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	registry.InitRegistry()
	reg := registry.GetRegistry()
	g := e.Group("/api/v1")
	g.GET("/liver", func(ctx echo.Context) error {
		return reg.VirtualLiverHandler.List(ctx)
	})
	// TODO: デバッグ用データ投入のために作ったので後でちゃんとレイヤードアーキテクチャ+DIPで作り直す
	g.POST("/liver", func(ectx echo.Context) error {
		ctx := ectx.Request().Context()

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
		entity := &model.VirtualLiver{
			Name:     "御伽原江良",
			NickName: "頭舞踏会",
		}

		key := client.IncompleteKey("virtual-liver", nil)
		_, err = client.Put(ctx, key, entity)
		if err != nil {
		}
		return ectx.JSON(http.StatusOK, entity)
	})
}
