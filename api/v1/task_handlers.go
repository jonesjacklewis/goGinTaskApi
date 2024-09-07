package v1

import (
	"database/sql"
	"task-list/controllers"

	"github.com/gin-gonic/gin"
)

// SetUpUserRoutes sets up the user routes
func SetUpTaskRoutes(db *sql.DB, router *gin.Engine) {

	router.GET("/tasks/user", controllers.GetTasksForUser)

	router.POST("/addTask", controllers.AddTask)
}
