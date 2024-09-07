package v1

import (
	"database/sql"
	"task-list/controllers"

	"github.com/gin-gonic/gin"
)

var shared_db *sql.DB

// SetUpUserRoutes sets up the user routes
func SetUpUserRoutes(db *sql.DB, router *gin.Engine) {
	shared_db = db

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/addUser", controllers.AddUser)
}
