package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdarify1337/backend-go/config"
	"github.com/mdarify1337/backend-go/modules"
)

func GetUsers(c *gin.Context) {
	var users []modules.User
	config.DB.Preload("Books").Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user modules.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user modules.User
	if err := config.DB.Preload("Books").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
