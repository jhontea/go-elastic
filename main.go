package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jhontea/go-elastic/docs"
	"github.com/jhontea/go-elastic/objects"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	elastic "gopkg.in/olivere/elastic.v6"
)

func main() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Elastic API for learning"
	docs.SwaggerInfo.Description = "This is an Elastic API for learning"
	docs.SwaggerInfo.Version = "1.0"

	router := gin.Default()
	port := ":3001"

	router.GET("v1/elastic/version", GetElasticVersion)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
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
func GetElasticVersion(c *gin.Context) {
	response := objects.Response{
		Code:   http.StatusInternalServerError,
		Status: "failed",
	}

	clientURL := "http://localhost:9200"
	client, err := elastic.NewClient(elastic.SetURL(clientURL))

	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	message := "Elasticsearch version " + esversion

	response.Code = http.StatusOK
	response.Message = message
	response.Status = "success"

	c.JSON(http.StatusOK, response)
}
