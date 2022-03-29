package controllers

import (
	"golang-gin-example/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{Name: "Luiz", Email: "example.com", Password: "123456"},
	{Name: "Eduardo", Email: "example.com", Password: "123456"},
	{Name: "Flavio", Email: "example.com", Password: "123456"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func PostUsers(c *gin.Context) {
	var newUsers models.User

	if err := c.BindJSON(&newUsers); err != nil {
		return
	}

	users = append(users, newUsers)
	c.IndentedJSON(http.StatusCreated, newUsers)
}
