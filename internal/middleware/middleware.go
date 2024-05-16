package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
)

//lint:ignore SA1029 "Using built-in type string as key for context value intentionally"
var ctx = context.WithValue(context.Background(), "packageName", "middleware")

func GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ins_log.SetUtfi("")
		startTime := time.Now() // Registro de inicio de tiempo
		// Aquí puedes realizar cualquier acción que deseas realizar antes de que se maneje la solicitud
		ins_log.Info(ctx, "New petition received")
		ins_log.Tracef(ctx, "url: %v, method: %v", c.Request.RequestURI, c.Request.Method)
		// Pasar la solicitud al siguiente middleware o al controlador final
		c.Next()
		elapsedTime := time.Since(startTime)
		ins_log.Infof(c.Request.Context(), "Request took %v", elapsedTime)
	}
}
