package services

import (
	"github.com/bijan/WarpFetch/internal/scraper"
)

// SearchResult represents a single search result
type SearchResult struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// SearchService performs the search operation using the scraper
func SearchService(query string) ([]SearchResult, error) {
	// Initialize the scraper
	scraper := scraper.NewScraper()

	// Perform the search
	results, err := scraper.Search(query)
	if err != nil {
		return nil, err
	}

	return results, nil
}
