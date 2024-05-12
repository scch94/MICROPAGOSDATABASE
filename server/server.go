package server

import (
	"context"
	"net/http"

	"github.com/scch94/MICROPAGOSDATABASE.git/internal/handler"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/routes"
	"github.com/scch94/ins_log"
)

const port = ":2121"

func StartServer() error {
	ctx := context.WithValue(context.Background(), "packageName", "server")
	h := &handler.Handler{}
	router := routes.SetupRouter(h)
	serverConfig := &http.Server{
		Addr:    port,
		Handler: router,
	}
	err := serverConfig.ListenAndServe()
	if err != nil {
		ins_log.Errorf(ctx, "cant connect to the server: %+v", err)
		return err
	}
	return nil
}
