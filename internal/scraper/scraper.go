package scraper

import (
	"github.com/bijan/WarpFetch/internal/models"
	"github.com/gocolly/colly"
)

// Scraper handles web scraping operations
type Scraper struct {
	collector *colly.Collector
}

// NewScraper creates a new Scraper instance
func NewScraper() *Scraper {
	c := colly.NewCollector(
		colly.AllowedDomains("www.google.com", "google.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	return &Scraper{
		collector: c,
	}
}

// Search performs a web search and returns the results
func (s *Scraper) Search(query string) ([]models.SearchResult, error) {
	var results []models.SearchResult

	s.collector.OnHTML("div.g", func(e *colly.HTMLElement) {
		title := e.ChildText("h3")
		url := e.ChildAttr("a", "href")
		description := e.ChildText("div.VwiC3b")

		if title != "" && url != "" {
			results = append(results, models.SearchResult{
				Title:       title,
				URL:         url,
				Description: description,
			})
		}
	})

	// Visit Google search results
	err := s.collector.Visit("https://www.google.com/search?q=" + query)
	if err != nil {
		return nil, err
	}

	return results, nil
}
