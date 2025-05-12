package services

import (
	"github.com/bijan/WarpFetch/internal/models"
	"github.com/bijan/WarpFetch/internal/scraper"
)

// SearchService performs the search operation using the scraper
func SearchService(query string) ([]models.SearchResult, error) {
	// Initialize the scraper
	s := scraper.NewScraper()

	// Perform the search
	results, err := s.Search(query)
	if err != nil {
		return nil, err
	}

	return results, nil
}
