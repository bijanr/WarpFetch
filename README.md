# WarpFetch

WarpFetch is a scalable web search API built with Go and the Gin framework. It provides a clean, modular architecture for web scraping and search functionality.

## Project Structure

```
WarpFetch/
├── cmd/                    # Main application entry point
├── internal/              # Internal packages
│   ├── handlers/         # HTTP request handlers
│   ├── services/         # Business logic
│   └── scraper/          # Web scraping implementation
├── api/                  # API-related code
│   └── routers/         # Route definitions
└── web/                 # Frontend demo
```

## Features

- RESTful API with Gin framework
- Web scraping using Colly
- Clean, modular architecture
- CORS support
- Simple frontend demo

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/bijan/WarpFetch.git
cd WarpFetch
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

1. Start the server:
```bash
go run cmd/main.go
```

2. Access the web interface:
Open your browser and navigate to `http://localhost:8080/web`

## API Endpoints

### Search
- **URL**: `/api/search`
- **Method**: `GET`
- **Query Parameters**:
  - `q`: Search query string
- **Response**:
```json
{
    "results": [
        {
            "title": "Result Title",
            "url": "https://example.com",
            "description": "Result description"
        }
    ]
}
```

## Development

The project follows a clean architecture pattern:

1. **Handlers** (`internal/handlers/`): Handle HTTP requests and responses
2. **Services** (`internal/services/`): Implement business logic
3. **Scraper** (`internal/scraper/`): Handle web scraping operations
4. **Routers** (`api/routers/`): Define API routes

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 