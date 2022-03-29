package main

import (
	"golang-gin-example/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	response := gin.Default()
	response.GET("/", controllers.GetUsers)
	response.POST("/", controllers.PostUsers)

	response.Run("localhost:5000")
}
