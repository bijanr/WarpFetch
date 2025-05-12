package routers

import (
	"github.com/bijan/WarpFetch/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine) {
	// API routes
	api := router.Group("/api")
	{
		api.GET("/search", handlers.SearchHandler)
	}

	// Serve static files from web directory
	router.Static("/web", "./web")
}
