package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhontea/go-elastic/docs"
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

func GetElasticVersion(c *gin.Context) {
	// Create a client and connect to http://192.168.2.10:9201
	clientURL := "http://localhost:9200"
	client, err := elastic.NewClient(elastic.SetURL(clientURL))
	if err != nil {
		// Handle error
	}
	fmt.Println(client, err)

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	message := "Elasticsearch version " + esversion

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
	})
}
