package main

import (
	"log"
	"database"
	"routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	db := database.Connect()
	defer db.Close()

	// Create a new gin router
	router := gin.Default()

	// Set up routes
	routes.Setup(router, db)

	// Start the server
	log.Fatal(router.Run(":8080"))
}	
