package server

import (
	"context"
	"net/http"

	"github.com/scch94/MICROPAGOSDATABASE.git/config"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/routes"
	"github.com/scch94/ins_log"
)

func StartServer(ctx context.Context) error {

	// Agregamos el valor "packageName" al contexto
	ctx = ins_log.SetPackageNameInContext(ctx, "server")

	ins_log.Infof(ctx, "Starting server on address: %s", config.Config.ServerPort)

	router := routes.SetupRouter(ctx)
	serverConfig := &http.Server{
		Addr:    config.Config.ServerPort,
		Handler: router,
	}
	err := serverConfig.ListenAndServe()
	if err != nil {
		ins_log.Errorf(ctx, "cant connect to the server: %+v", err)
		return err
	}
	return nil
}
