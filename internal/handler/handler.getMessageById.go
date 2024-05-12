package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scch94/MICROPAGOSDATABASE.git/database"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/responses"
	"github.com/scch94/ins_log"
)

func (handler *Handler) GetMessageById(c *gin.Context) {
	//getting the id and the utfi in the params
	idStr := c.Param("id")
	//parcing the id to int 64
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ins_log.Error(ctx, "error parsing de id in a Uint:")
		// Manejar el error si la conversión falla
		// Por ejemplo, devolver un error al cliente
		response := responses.NewResponseMessage(1, err.Error(), 0)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	utfi := c.Param("utfi")
	if utfi == "" {
		ins_log.GenerateUtfi()
	} else {
		ins_log.SetUtfi(utfi)
	}
	ins_log.Infof(ctx, "starting to get the message with id %s", idStr)
	//getting the pool conection to insert
	ins_log.Tracef(ctx, "startin to get the pool conection to get the message")
	storageMessage := database.NewMysqlMessage(database.PoolMessage())
	serviceMessage := models.NewService(storageMessage)
	//creamos una variable con el modelo del message

	modelMessage, err := serviceMessage.GetByID(id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		ins_log.Infof(ctx, "product with id: %s not found", idStr)
		message := "this id is not register in the database " + idStr
		response := responses.NewResponseMessage(0, message, modelMessage.Id)
		c.JSON(http.StatusOK, response)
		return
	case err != nil:
		ins_log.Error(ctx, "error getting the message by id :")
		response := responses.NewResponseMessage(1, err.Error(), 0)
		c.JSON(http.StatusInternalServerError, response)
		return
	default:
		ins_log.Infof(ctx, "the message with id : %s is this: %s", idStr, modelMessage.Content)
	}
	message := "this is the message " + modelMessage.Content
	response := responses.NewResponseMessage(0, message, id)
	c.JSON(http.StatusOK, response)
}
