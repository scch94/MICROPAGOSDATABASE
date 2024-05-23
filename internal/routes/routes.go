package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/handler"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/middleware"
	"github.com/scch94/ins_log"
)

func SetupRouter(ctx context.Context) *gin.Engine {

	// create a new gin router and register the handlers
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Agregar middleware global
	router.Use(gin.Recovery())
	router.Use(middleware.GlobalMiddleware())

	h := handler.Handler{}

	//rutas
	router.GET("/", h.Welcome)
	router.POST("/insertMessage/:utfi", h.InsertMessage)
	router.GET("/filter/:mobile/:shortNumber/:utfi", h.IsFilter)
	router.GET("/message/:id/:utfi", h.GetMessageById)
	router.GET("/userdomain/:username/:utfi", h.GetUserDomain)
	router.NoRoute(notFoundHandler)

	return router
}

// Controlador para manejar rutas no encontradas
func notFoundHandler(c *gin.Context) {
	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")
	ins_log.Errorf(ctx, "Route  not found: url: %v, method: %v", c.Request.RequestURI, c.Request.Method)
	c.JSON(http.StatusNotFound, nil)
}
