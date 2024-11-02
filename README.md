# go-gorm-mk1-showcase

A showcase project demonstrating usage of [GORM](https://gorm.io/) with Go, focusing on essential database operations within a simple web server built with [Fiber](https://gofiber.io/).

This project implements CRUD operations on a PostgreSQL database, covering fundamental usage patterns for GORM in Go, and provides a lightweight API interface with basic HTML templating.

## Project Structure

```
go-gorm-mk1-showcase/
├── gorm/
│   └── gorm.go
├── server/
│   ├── handlers/
│   │   └── handlers.go
│   └── templates/
│       └── index.html
├── go.mod
├── go.sum
└── main.go
```

## Code Overview

**Main Components**

- main.go: The entry point that initializes GORM, sets up the Fiber application, configures routes, and starts the server.
- gorm/gorm.go: Contains GORM configuration, database connection, and model definition. Provides functions to perform CRUD operations on the purchases table.
- handlers/handlers.go: Defines route handlers for each endpoint, mapping HTTP requests to corresponding GORM functions and handling error responses.
- templates/index.html: Basic HTML template used to render the root endpoint.

**GORM Model**

The Purchase model represents a database entity with fields for ID, Name, Description, and Amount:

```go
type Purchase struct {
    ID          uint   `gorm:"type:BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT"`
    Name        string `gorm:"type:VARCHAR(100) NOT NULL"`
    Description string `gorm:"type:VARCHAR(150)"`
    Amount      int    `gorm:"type:INT NOT NULL"`
}
```

**Database Initialization**

The ConfigGorm function sets up the database connection and ensures the purchases table exists. If the table doesn’t exist, it creates it and populates it with initial data.

**Error Handling**

The custom Error handler in handlers/handlers.go is used to manage errors throughout the application:
- Returns a 404 message if the requested page is not found.
- Displays a generic error message for all other issues.

## Requirements

- **Go** (v1.16 or later)
- **PostgreSQL** (Ensure the database server is running and accessible)

## Setup and Configuration

### 1. Clone the Repository

```bash
git clone https://github.com/xoticdsign/go-gorm-mk1-showcase
```

```bash
cd go-gorm-mk1-showcase
```

### 2. Configure Database Connection

In gorm/gorm.go, configure the database connection parameters:

```go
var (
	host     string // Your host (example: localhost)
	user     string // Your user (example: postgres)
	password string // Your password (optional)
	dbname   string // Your DB name (example: postgres)
	sslmode  string // SSL (enable/disable)
)
```

Update these values based on your PostgreSQL setup.

### 3. Run the Application

```bash
go run main.go
```

The server will start on http://0.0.0.0:3240 and will be accessible on any device in your WI-FI network.

## Usage

The application exposes several routes to interact with the database. Here are the main routes and their functions:

**Root Endpoint**

```
GET /: Renders the index.html template.
```

**CRUD Endpoints**

Each endpoint performs a specific database operation using GORM:

```
GET /select-all: Retrieves all records from the purchases table.
GET /select-where: Retrieves a record with name="PS5" from the purchases table.
GET /select-specific: Selects specific fields (name, amount) for all records.
GET /update-all: Updates the record with name="PS5" to new values and returns the updated record.
GET /update-name: Updates the name field for a specific record (name="Sneakers") to “My Purchase.”
GET /delete-row: Deletes the record with name="Grocery" and returns the count of affected rows.
GET /insert-row: Inserts a new record into the database and returns the inserted record.
```

## License

[MIT]()