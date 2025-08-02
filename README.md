
# Go LINQ API

This project provides a RESTful API for querying Vietnamese administrative units (provinces, wards, etc.) using Go, GORM, and Gin. It demonstrates LINQ-like query building and modular service/repository/controller architecture.

## Getting Started

### Prerequisites

- Go 1.24+
- SQL Server database with the `vietnamese_administrative_units` schema

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/go-linq-api.git
   cd go-linq-api
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Configure your database connection in `internal/db/connection.go`.

### Running the Project

Start the API server:
```sh
go run cmd/main.go
```
The server runs at `http://localhost:8080`.

## Project Structure

- **cmd/main.go**: Entry point. Initializes DB, repositories, services, controllers, and starts Gin server.
- **internal/db/connection.go**: Database connection logic.
- **internal/models/**: Data models for provinces, wards, administrative units, and regions.
- **internal/repositories/**: Data access layer. Implements queries using GORM and LINQ-like builder (`internal/linq/linq.go`).
- **internal/services/**: Business logic layer. Calls repository methods.
- **internal/controllers/**: HTTP handlers. Maps API routes to service methods.
- **internal/helpers/pagination-utilities.go**: Utilities for paginating results.

## API Endpoints

### Provinces

- `GET /api/provinces/`  
  List all provinces.

- `GET /api/provinces/:code`  
  Get province by code.

- `GET /api/provinces/stats`  
  Get province statistics (with ward count and unit name).

### Wards

- `GET /api/wards/`  
  List all wards.

- `GET /api/wards/details`  
  Get ward details (with province and unit name).

## Custom LINQ-like Query Builder

See `internal/linq/linq.go` for a flexible query builder supporting joins, selects, grouping, ordering, and more.

## Pagination

Utilities for paginating results are in `internal/helpers/pagination-