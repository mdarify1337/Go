package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mdarify1337/backend-go/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	r.POST("/users/:id/books", controllers.AddBookToUser)
	r.GET("/users/:id/books", controllers.GetUserBooks)
}
