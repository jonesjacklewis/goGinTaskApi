package controllers

import (
	"fmt"
	"net/http"
	"task-list/repositories"

	"github.com/gin-gonic/gin"
)

// GetAllUsers controller to get all users
func GetAllUsers(c *gin.Context) {
	users, err := repositories.UserRepository.GetAllUsers()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to get users"})
		fmt.Println("Failed to get users:", err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"error": false, "data": gin.H{"Users": users}})

}

// AddUser controller to add a user
func AddUser(c *gin.Context) {
	var display_name string = c.PostForm("display_name")

	if display_name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": true, "message": "display_name is required"})
		return
	}

	user, err := repositories.UserRepository.CreateUser(display_name)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Failed to create user"})
		fmt.Println("Failed to create user:", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"error": false, "data": gin.H{
		"User": user,
	}})
}
