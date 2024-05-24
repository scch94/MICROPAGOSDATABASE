package main

import (
	"context"
	"io"
	"os"
	"strings"
	"time"

	"github.com/scch94/MICROPAGOSDATABASE.git/config"
	"github.com/scch94/MICROPAGOSDATABASE.git/database"
	"github.com/scch94/MICROPAGOSDATABASE.git/server"
	"github.com/scch94/ins_log"
)

func main() {
	// Creamos el contexto para esta ejecución
	ctx := context.Background()
	// Obtener la fecha actual

	today := time.Now().Format("2006-01-02 15")
	// Reemplazar los caracteres no permitidos en el nombre del archivo
	replacer := strings.NewReplacer(" ", "-")
	today = replacer.Replace(today)

	// Construir el nombre del archivo de log
	logFileName := "micropagosdatabasegateway_" + today + ".log"
	file, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Creamos un escritor que escriba tanto en el archivo como en la consola
	multiWriter := io.MultiWriter(os.Stdout, file)
	ins_log.StartLoggerWithWriter(multiWriter)

	//levantamos la config
	errConfig := config.Upconfig(ctx)
	if errConfig != nil {
		ins_log.Errorf(ctx, "error when we try to get the configuration err: %v", errConfig)
		return
	}

	// Inicializamos el logger
	ins_log.SetService("micropagosdatabase")
	ins_log.SetLevel(config.Config.LogLevel)

	// Agregamos el valor "packageName" al contexto
	ctx = ins_log.SetPackageNameInContext(ctx, "main")

	ins_log.Infof(ctx, "startig micropagos database module version : %+v", version())

	//conectando a las 2 bases de datos
	driverUserDatabase := database.MySQLUsers
	database.New(driverUserDatabase, ctx)
	driverMysql := database.MySQL
	database.New(driverMysql, ctx)

	err = server.StartServer(ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error al tratarde iniciar el servidor : %s", err.Error())
	}
}

func version() string {
	return "1.0.0"
}
