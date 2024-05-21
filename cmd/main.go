package main

import (
	"context"
	"os"

	"github.com/scch94/MICROPAGOSDATABASE.git/config"
	"github.com/scch94/MICROPAGOSDATABASE.git/database"
	"github.com/scch94/MICROPAGOSDATABASE.git/server"
	"github.com/scch94/ins_log"
)

//lint:ignore SA1029 "Using built-in type string as key for context value intentionally"
var ctx = context.WithValue(context.Background(), "packageName", "main")

func main() {
	//levantamos la config
	errConfig := config.Upconfig()
	if errConfig != nil {
		ins_log.Errorf(ctx, "error when we try to get the configuration err: %v", errConfig)
		return
	}

	//inicialisamos el logger
	ins_log.StartLogger()
	ins_log.SetService("micropagosdatabase")
	ins_log.Infof(ctx, "startig micropagos database module version : %+v", version())

	//conectando a las 2 bases de datos
	driverUserDatabase := database.MySQLUsers
	database.New(driverUserDatabase)
	driverMysql := database.MySQL
	database.New(driverMysql)

	// Abrir el archivo de log
	file, err := os.OpenFile("logfile.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		ins_log.Errorf(ctx, "Error al abrir el archivo de log: %s", err.Error())
		return
	}
	defer file.Close()
	/*
		// Crear un writer multi para enviar logs tanto al archivo como a la consola
		multiWriter := io.MultiWriter(os.Stdout, file)

		// Inicializar el logger con el writer multi
		ins_log.StartLoggerWithWriter(multiWriter)
	*/
	//inicamos el servidor
	err = server.StartServer()
	if err != nil {
		ins_log.Errorf(ctx, "error al tratarde iniciar el servidor : %s", err.Error())
	}
}

func version() string {
	return "1.0.0"
}
