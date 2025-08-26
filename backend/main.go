package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdarify1337/backend-go/config"
	"github.com/mdarify1337/backend-go/modules"
	"github.com/mdarify1337/backend-go/routes"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&modules.User{}, &modules.Book{})
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
