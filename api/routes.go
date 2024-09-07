package api

import (
	v1 "task-list/api/v1"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1.SetUpUserRoutes(nil, router)
	v1.SetUpTaskRoutes(nil, router)

	return router
}
