package main

import (
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/jhontea/go-elastic/connections"
	elasticController "github.com/jhontea/go-elastic/controllers/elastics"
)

func Route() {
	router := gin.Default()
	port := os.Getenv("PORT")

	client := *connections.ElasticInit()
	controller := elasticController.V1ElasticControllerHandler(client)
	router.GET("v1/elastic/version", controller.GetElasticVersion)
	router.GET("v1/elastic/field/detail", controller.GetByField)
	router.GET("v1/elastic/detail", controller.GetById)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
}
