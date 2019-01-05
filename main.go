package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := ":3001"

	router.Run(port)
}
