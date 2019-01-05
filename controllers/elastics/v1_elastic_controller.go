package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/jhontea/go-elastic/services/elastics"
	elastic "gopkg.in/olivere/elastic.v6"
)

type V1ElasticController struct {
	ElasticService services.V1ElasticService
}

func V1ElasticControllerHandler(client elastic.Client) V1ElasticController {
	controller := &V1ElasticController{
		ElasticService: services.V1ElasticServiceHandler(client),
	}

	return *controller
}

// GetElasticVersion godoc
// @Summary Elasticsearch version
// @Description Show Elasticsearch version from url
// @Tags Elastic
// @Accept  json
// @Produce  json
// @Success 200 {object} objects.Response
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/elastic/version [get]
func (controller *V1ElasticController) GetElasticVersion(c *gin.Context) {
	response, err := controller.ElasticService.GetElasticVersion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
