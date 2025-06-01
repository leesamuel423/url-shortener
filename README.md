# URL Shortener

URL shortening service built with Go. Allows users to create short URLs, track click statistics, and manage custom short codes.

![demo](./demo.gif)

## Prerequisites

- Go 1.23 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (optional)
- Make (optional, for using Makefile commands)

## Installation

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

2. Set up environment variables:
```bash
cp .env.example .env
```

Edit `.env` with your configuration:
```env
DATABASE_URL=postgres://postgres:password@localhost:5432/urlshortener?sslmode=disable
DOMAIN=http://localhost:8080
PORT=8080
ENV=development
```

3. Install dependencies:
```bash
make deps
```

4. Set up the database:
```bash
# Create the database
createdb urlshortener

# Run migrations (handled automatically on startup)
```

5. Run the application:
```bash
make run
```

The application will be available at `http://localhost:8080`

### Docker Deployment

1. Build and run with Docker Compose:
```bash
make docker-up
```

2. Stop the containers:
```bash
make docker-down
```

## Makefile Commands

### Build & Run
- `make build` - Build the application binary
- `make run` - Run the application
- `make dev` - Run with hot reload (requires Air)
- `make clean` - Clean build artifacts

### Testing & Quality
- `make test` - Run tests
- `make test-coverage` - Run tests with coverage report
- `make lint` - Run linter (requires golangci-lint)
- `make fmt` - Format code

### Dependencies
- `make deps` - Download dependencies
- `make tidy` - Tidy and verify dependencies

### Docker
- `make docker-build` - Build Docker image
- `make docker-up` - Start Docker containers
- `make docker-down` - Stop Docker containers
- `make docker-logs` - View container logs

### Database
- `make migrate-up` - Run database migrations
- `make migrate-down` - Rollback migrations
- `make migrate-create name=<name>` - Create new migration

## Development

### Running Tests
```bash
make test
```

### Running with Hot Reload
```bash
# Install Air first
go install github.com/air-verse/air@latest

# Run with hot reload
make dev
```

### Code Formatting
```bash
make fmt
```

### Linting
```bash
# Install golangci-lint first
brew install golangci-lint  # macOS
# or
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

make lint
```

## API Endpoints

### Shorten URL
```http
POST /api/shorten
Content-Type: application/json

{
  "url": "https://example.com",
  "custom_code": "optional-custom-code"
}

Response:
{
  "short_url": "http://localhost:8080/abc123",
  "original_url": "https://example.com"
}
```

### Get URL Statistics
```http
GET /api/stats/{shortCode}

Response:
{
  "url": {
    "id": 1,
    "short_code": "abc123",
    "original_url": "https://example.com",
    "created_at": "2025-05-31T12:00:00Z",
    "click_count": 42
  },
  "short_url": "http://localhost:8080/abc123"
}
```

### Redirect
```http
GET /{shortCode}

Response: 301 Redirect to original URL
```

## Configuration

The application can be configured through environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgres://postgres:password@localhost/urlshortener?sslmode=disable` |
| `DOMAIN` | Base domain for shortened URLs | `http://localhost:8080` |
| `PORT` | Server port | `8080` |
| `ENV` | Environment (development/production) | `development` |


## Database Schema

The application uses a single `urls` table:

```sql
CREATE TABLE urls (
    id BIGSERIAL PRIMARY KEY,
    short_code VARCHAR(10) UNIQUE NOT NULL,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    click_count INTEGER DEFAULT 0
);

CREATE INDEX idx_short_code ON urls(short_code);
CREATE INDEX idx_original_url ON urls(original_url);
```
