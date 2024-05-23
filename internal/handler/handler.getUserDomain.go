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

func (handler *Handler) GetUserDomain(c *gin.Context) {

	ctx := c.Request.Context()
	ctx = ins_log.SetPackageNameInContext(ctx, "handler")

	request := request.GetUserDomain{
		UserName: c.Param("username"),
	}

	ins_log.Infof(ctx, "starting to get the domain for the user %s", request.UserName)

	ins_log.Tracef(ctx, "startin to get the pool conection to get the message")
	storageDomain := database.NewMysqlDomain(database.PoolUsers())
	serviceDomain := models.NewDomainService(storageDomain)

	//REALISAMOS LA CONSULTA
	domainResponse, err := serviceDomain.GetUserDomain(request, ctx)
	if err != nil {
		ins_log.Errorf(ctx, "error getting Domain response")
		response := responses.NewResponseMessage(1, err.Error(), 0)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	//creamos la respuesta
	response := responses.DomainResponse{
		Response: responses.Response{
			Result:  0,
			Message: domainResponse.Result,
		},
		DomainName: domainResponse.Domainname,
	}
	c.JSON(http.StatusOK, response)

}
