package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdarify1337/backend-go/config"
	"github.com/mdarify1337/backend-go/modules"
)

func GetBooks(c *gin.Context) {
	var books []modules.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book modules.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func AddBookToUser(c *gin.Context) {
	userID := c.Param("id")
	uid, _ := strconv.ParseUint(userID, 10, 32)

	var book modules.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.UserID = uint(uid)
	config.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func GetUserBooks(c *gin.Context) {
	userID := c.Param("id")
	var books []modules.Book
	config.DB.Where("user_id = ?", userID).Find(&books)
	c.JSON(http.StatusOK, books)
}
