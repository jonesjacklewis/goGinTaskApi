package v1

import (
	"database/sql"
	"task-list/controllers"

	"github.com/gin-gonic/gin"
)

// SetUpUserRoutes sets up the user routes
func SetUpUserRoutes(db *sql.DB, router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers)
	router.POST("/addUser", controllers.AddUser)
}
