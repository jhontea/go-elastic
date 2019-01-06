package main

import (
	"github.com/jhontea/go-elastic/docs"
	"github.com/joho/godotenv"
)

func main() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Elastic API for learning"
	docs.SwaggerInfo.Description = "This is an Elastic API for learning"
	docs.SwaggerInfo.Version = "1.0"

	// load .env
	godotenv.Load()

	Route()
}
