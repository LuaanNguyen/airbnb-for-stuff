package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB is the database connection pool
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
    // Load .env file explicitly
    err := godotenv.Load(".env")
    if err != nil {
        return fmt.Errorf("error loading .env file: %v", err)
    }

    fmt.Println("Connecting to Postgres...")
    
    // Use the POSTGRES_URL directly
    connStr := os.Getenv("POSTGRES_URL")
    
    // Open database connection
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error opening database: %v", err)
    }

    // Test the connection
    err = DB.Ping()
    if err != nil {
        return fmt.Errorf("error connecting to database: %v", err)
    }

    fmt.Println("Successfully connected to database!")
    return nil
}