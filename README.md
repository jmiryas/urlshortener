# URL Shortener Service

A RESTful URL shortener service built with Go, Fiber, GORM, and PostgreSQL that converts long URLs into short, easily shareable links with analytics tracking.

## Features

- URL shortening with auto generated tokens
- Redirect to original URLs
- Click analytics tracking (unique visitors, referrer)
- User authentication using JWT tokens
- Idempotent token generation (same URL = same token)
- Docker containerization for easy deployment
- Comprehensive test coverage

## Technology Stack

- **Backend**: Go 1.21+ with Fiber framework
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT tokens
- **Containerization**: Docker and Docker Compose
- **Testing**: Native Go testing framework

## Project Structure

```
urlshortener/
├── config/
│   └── config.go
├── handlers/
│   ├── analytics.go
│   ├── auth.go
│   ├── redirect.go
│   ├── shorten.go
│   └── stats.go
├── migrations/
│   ├── 001_init_schema.up.sql
│   └── 001_init_schema.down.sql
├── models/
│   ├── url.go
│   ├── user.go
│   └── visit.go
├── routes/
│   └── routes.go
├── storage/
│   ├── database.go
│   └── migrate.go
├── tests/
│   ├── helpers_test.go
│   ├── redirect_test.go
│   ├── shorten_test.go
│   └── stats_test.go
├── utils/
│   ├── token.go
│   └── url.go
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── Makefile
```

## API Endpoints

1. Shorten URL

   Endpoint: `POST /api/v1/shorten`

   Creates a short URL from a long URL.

   Request:

   ```bash
   curl -X POST -H "Content-Type: application/json" \
   -d '{"url":"https://example.com/very/long/url/path"}' \
   http://localhost:3000/shorten
   ```

   Response:

   ```json
   {
     "token": "abc123",
     "short_url": "http://localhost:3000/abc123",
     "long_url": "https://example.com/very/long/url/path"
   }
   ```

2. Redirect to Original URL

   Endpoint: `GET /api/v1/:token`

   Redirects to the original URL.

   Request:

   ```bash
   curl -I http://localhost:3000/api/v1/abc123
   ```

   Response:

   ```bash
   HTTP/1.1 302 Found
   Location: https://example.com/very/long/url/path
   Date: Wed, 03 Sep 2025 10:00:00 GMT
   ```

3. Get URL Statistics

   Endpoint: `GET /api/v1/stats/:token`

   Retrieves analytics for a short URL.

   Request:

   ```bash
   curl http://localhost:3000/api/v1/stats/abc123
   ```

   Response:

   ```json
   {
     "click_count": 2,
     "created_at": "2025-09-03T00:46:12.478646Z",
     "original_url": "https://example.com/very/long/url/path",
     "short_token": "abc123"
   }
   ```

4. Analytics

   Endpoint: `POST /api/v1/analytics/:token`

   Retrieves analytics for a short URL.

   Request:

   ```bash
   curl http://127.0.0.1:3000/api/v1/analytics/abc123
   ```

   Response:

   ```json
   {
     "referrers": [
       {
         "referrer": "http://localhost:5000/",
         "count": 9
       }
     ],
     "total_clicks": 11,
     "unique_visitors": 1,
     "url": {
       "click_count": 11,
       "original_url": "http://example.com/",
       "short_token": "abc123"
     }
   }
   ```

5. User Registration

   Endpoint: `POST /api/v1/auth/register`

   Creates a new user account.

   Request:

   ```bash
   curl -X POST -H "Content-Type: application/json" \
   -d '{"name": "Budi", "username":"budi","password":"123456"}' \ http://localhost:3000/api/v1/auth/register
   ```

   Response:

   ```json
   {
     "message": "User created successfully",
     "user": {
       "id": 1,
       "name": "Budi",
       "username": "budi"
     }
   }
   ```

6. User Login

   Endpoint: `POST /api/v1/auth/login`

   Authenticates a user and returns a JWT token.

   Request:

   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"username":"budi","password":"123456"}' http://localhost:3000/api/v1/auth/login
   ```

   Response:

   ```json
   {
     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
     "user": {
       "id": 1,
       "username": "budi"
     }
   }
   ```

## Setup Instructions

### Prerequisites

- Go 1.21+
- PostgreSQL
- Docker (optional)

### Local Development

1. Clone the repository:

   ```bash
   git clone https://github.com/jmiryas/urlshortener.git
   cd urlshortener
   ```

1. Set up environment variables:

   ```bash
   cp .env.example .env
   ```

1. Install dependencies:

   ```bash
   go mod download
   ```

1. Start the server:

   ```bash
   # if using docker

   docker-compose build --no-cache app

   docker-compose up

   # without docker

   go run main.go
   ```

### Docker Deployment

1. Build and start all containers:

   ```bash
   docker-compose down

   docker-compose build --no-cache app

   docker-compose up --build
   ```

1. The service will be available at `http://localhost:3000`

## Environment Variables

| Variable      | Deskripsi                | Default |
| ------------- | ------------------------ | ------- |
| `PORT`        | Server port              | `3000`  |
| `DB_HOST`     | Database host            | -       |
| `DB_USER`     | Database user            | -       |
| `DB_PASSWORD` | Database password        | -       |
| `DB_NAME`     | Database name            | -       |
| `DB_PORT`     | Database port            | `5432`  |
| `JWT_SECRET`  | Secret untuk JWT signing | -       |

## Running Tests

Execute the test suite with:

```bash
go test ./tests/ -v

# or

make test
```

Result:

```bash
=== RUN   TestRedirect
=== RUN   TestRedirect/Positive:_valid_redirect
=== RUN   TestRedirect/Negative:_unknown_token
--- PASS: TestRedirect (0.01s)
    --- PASS: TestRedirect/Positive:_valid_redirect (0.00s)
    --- PASS: TestRedirect/Negative:_unknown_token (0.00s)
...

# or

PASS tests.TestRedirect/Positive:_valid_redirect (0.01s)
PASS tests.TestRedirect/Negative:_unknown_token (0.00s)
...
```

## Test Cases Coverage

The implementation covers all required test cases:

### TC-1: Generate Short URL

- ✅ Valid URL returns 201 with short token
- ✅ Missing URL returns 400 Bad Request
- ✅ Invalid URL format returns 422 Unprocessable Entity

### TC-2: Redirect to Long URL

- ✅ Valid token redirects with 302 Found
- ✅ Unknown token returns 404 Not Found

### TC-3: Idempotent Token Generation

- ✅ Same URL generates identical token

### TC-4: Click Analytics Tracking

- ✅ Click counter increments on each access
- ✅ Stats endpoint returns accurate count
- ✅ Invalid token returns 404 Not Found

## Design Rationale

### Architecture Decisions

- Used Fiber for its Express-like simplicity and performance
- Implemented GORM for database abstraction and migrations
- Chose PostgreSQL for reliability and JSON support
- Used JWT for stateless authentication

### Token Generation

- Uses base62 encoding of database ID for short URLs
- Ensures uniqueness through database constraints
- Provides predictable length based on record count

### Idempotency

- Same URL always generates same token by checking existing records
- Prevents duplicate entries for the same URL

### Analytics

- Tracks each visit with timestamp
- Provides click count and last accessed time
- Easily extendable for more detailed analytics
