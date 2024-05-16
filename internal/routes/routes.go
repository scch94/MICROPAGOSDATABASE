package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/handler"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/middleware"
	"github.com/scch94/ins_log"
)

//lint:ignore SA1029 "Using built-in type string as key for context value intentionally"

var ctx = context.WithValue(context.Background(), "packageName", "routes")

func SetupRouter(h *handler.Handler) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	// create a new gin router and register the handlers
	router := gin.New()
	// Agregar middleware global
	router.Use(middleware.GlobalMiddleware())
	//globalMiddleware())
	// Middleware de recuperación para manejar errores de pá
	router.Use(gin.Recovery())
	router.GET("/", h.Welcome)
	router.POST("/insertMessage", h.InseretMessage)
	router.GET("/message/:id/:utfi", h.GetMessageById)
	router.NoRoute(notFoundHandler)

	return router
}

// isarasola / 4Gpeperoni5f2l&6519^P$$a
// Controlador para manejar rutas no encontradas
func notFoundHandler(c *gin.Context) {

	ins_log.Errorf(ctx, "Route  not found: url: %v, method: %v", c.Request.RequestURI, c.Request.Method)
	c.JSON(http.StatusNotFound, nil)
}
