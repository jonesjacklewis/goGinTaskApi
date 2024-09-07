package main

import (
	"fmt"
	"task-list/api"
	"task-list/config"
)

// main function
func main() {
	// Initialize the database connection
	err := config.InitDb()
	if err != nil {
		fmt.Println("Failed to initialize the database:", err)
	}
	defer config.CloseDb()

	// Setup the router with all routes
	router := api.SetupRouter()

	// Run the Gin router
	if err := router.Run(":8123"); err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
