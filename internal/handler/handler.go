package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/ins_log"
)

//lint:ignore SA1029 "Using built-in type string as key for context value intentionally"
var ctx = context.WithValue(context.Background(), "packageName", "handler")

type Handler struct {
}

func (h *Handler) Welcome(c *gin.Context) {

	ins_log.Info(ctx, "starting handler welcome")

	c.JSON(http.StatusOK, "bienvenidos")

}