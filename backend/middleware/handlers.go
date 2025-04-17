package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type Response struct {
	ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}


// create a DB connection 
func CreateConnections() *sql.DB {
	fmt.Println("Connecting to Postgres...")
	
	// load .env file 
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// open db connection 
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	// ping the DB
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
    // return the connection
    return db
}
