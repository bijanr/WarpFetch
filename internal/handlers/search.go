package handlers

import (
	"net/http"

	"github.com/bijan/WarpFetch/internal/services"
	"github.com/gin-gonic/gin"
)

// SearchHandler handles the search endpoint
func SearchHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'q' is required",
		})
		return
	}

	// Get search results from the service layer
	results, err := services.SearchService(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to perform search",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}
