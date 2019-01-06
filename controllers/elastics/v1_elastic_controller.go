package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	elastic "gopkg.in/olivere/elastic.v6"

	services "github.com/jhontea/go-elastic/services/elastics"
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
// @Accept json
// @Produce json
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

// GetByField godoc
// @Summary Get data by field
// @Description Show data by field from some index with value
// @Tags Elastic
// @Accept json
// @Produce json
// @Param field query string false "field"
// @Param value query string false "value"
// @Param index query string false "index"
// @Success 200 {object} objects.Response
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/elastic/field/detail [get]
func (controller *V1ElasticController) GetByField(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	index := c.Query("index")

	response := controller.ElasticService.GetByField(field, value, index)

	c.JSON(200, response)
}

// GetById godoc
// @Summary Get data by id
// @Description Show data by id from some index
// @Tags Elastic
// @Accept json
// @Produce json
// @Param uuid query string false "uuid"
// @Param index query string false "index"
// @Success 200 {object} objects.Response
// @Failure 500 {string} string "Internal Server Error"
// @Router /v1/elastic/detail [get]
func (controller *V1ElasticController) GetById(c *gin.Context) {
	uuid := c.Query("uuid")
	index := c.Query("index")

	response := controller.ElasticService.GetById(uuid, index)

	c.JSON(200, response)
}
