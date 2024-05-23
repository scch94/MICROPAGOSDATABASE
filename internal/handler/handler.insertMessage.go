package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/scch94/MICROPAGOSDATABASE.git/database"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models/responses"
	"github.com/scch94/ins_log"
)

func (handler *Handler) InsertMessage(c *gin.Context) {

	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")
	var request models.MessageModel
	// Utiliza el método BindJSON de Gin para vincular los datos del cuerpo de la solicitud a la estructura request
	if err := c.BindJSON(&request); err != nil {
		{
			ins_log.Errorf(ctx, "error when we try to get the json petition")

			response := responses.NewResponseMessage(1, err.Error(), 0)
			c.JSON(http.StatusBadRequest, response)
			return
		}

	}
	ins_log.Tracef(ctx, "this is the data that we recibed in the petition to insert the message %s", request)
	ins_log.Info(ctx, "starting to insert message")
	// Obtener la conexión de la pool para insertar
	storageMessage := database.NewMysqlMessage(database.PoolMessage())
	serviceMessage := models.NewService(storageMessage)

	err := serviceMessage.InsertMessage(&request, ctx)
	if err != nil {
		ins_log.Error(ctx, "error inserting message insertMessage(request) :")
		response := responses.NewResponseMessage(1, err.Error(), 0)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	message := "el mensaje se inserto correctamente con el id" + strconv.FormatUint(request.Id, 10)
	response := responses.NewResponseMessage(0, message, request.Id)
	c.JSON(http.StatusOK, response)
}
