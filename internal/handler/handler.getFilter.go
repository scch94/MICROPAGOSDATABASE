package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/MICROPAGOSDATABASE.git/database"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models/request"
	"github.com/scch94/MICROPAGOSDATABASE.git/internal/models/responses"
	"github.com/scch94/ins_log"
)

func (handler *Handler) IsFilter(c *gin.Context) {

	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	//creamos la structura donde vamos a guardar la info
	request := request.IsFiler{
		ShortNumber: c.Param("shortNumber"),
		Mobile:      c.Param("mobile"),
	}

	ins_log.Infof(ctx, "starting to check if the shortnumber %s with the mobile %s are filter", request.ShortNumber, request.Mobile)

	//getting the pool conection to filter
	ins_log.Tracef(ctx, "startin to get the pool conection to get the message")
	storageFilter := database.NewMysqlFilter(database.PoolMessage())
	serviceFiler := models.NewFilterService(storageFilter)

	//REALISAMOS LA CONSULTA
	filterResponse, err := serviceFiler.IsFilter(request, ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error getting filter response")
		response := responses.NewResponseMessage(1, err.Error(), 0)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//creamos la respuesta
	response := responses.IsFilterResponse{
		Result:  filterResponse.Result,
		Message: filterResponse.Comment,
	}
	c.JSON(http.StatusOK, response)

}
