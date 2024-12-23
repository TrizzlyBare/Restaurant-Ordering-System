package main

import (
	"log"

	"github.com/TrizzlyBare/Restaurant-Ordering-System/database"
	"github.com/TrizzlyBare/Restaurant-Ordering-System/models"
	"github.com/TrizzlyBare/Restaurant-Ordering-System/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	database.DB.AutoMigrate(&models.User{}, &models.Menu{}, &models.Order{}, &models.Payment{})

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.MenuRoutes(r)
	routes.OrderRoutes(r)

	log.Println("Server running on http://localhost:8080")
	r.Run(":8080")
}
