package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type Response struct {
	ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

//------------------------- HTTP functions ----------------

// create a DB connection 
func CreateConnection() *sql.DB {
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

// Check health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	db := CreateConnection() // establish connection with 
	defer db.Close()

	resp := Response {
		Message: "Hello, you have successfully connected to Postgres 🫶",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


// Get all users 
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the users in the db 
	users, err := getAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve all users", http.StatusInternalServerError)
		return
	}

	// send all users as response
	json.NewEncoder(w).Encode(users)
}

// Get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	user, err := getUser(int64(id))
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
	}

	// send user with matching id
	json.NewEncoder(w).Encode(user)
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	items, err := getAllItems()
	if err != nil {
		http.Error(w, "Failed to retrieve all items", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}