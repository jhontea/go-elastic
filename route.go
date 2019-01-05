package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jhontea/go-elastic/connections"
	elasticController "github.com/jhontea/go-elastic/controllers/elastics"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Route() {
	router := gin.Default()
	port := ":3001"

	client := *connections.ElasticInit()
	controller := elasticController.V1ElasticControllerHandler(client)

	router.GET("v1/elastic/version", controller.GetElasticVersion)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
}
