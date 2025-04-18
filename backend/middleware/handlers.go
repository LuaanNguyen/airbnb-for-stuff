package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LuaanNguyen/backend/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type Response struct {
	ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}


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


// Get all users 
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the users in the db 
	users, err := getAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	// send all users as response
	json.NewEncoder(w).Encode(users)
}



//------------------------- handler functions ----------------
func getAllUsers() ([]models.User, error) {
	// create the db connection 
	db := CreateConnection() 
	defer db.Close()
	
	sqlStatement := `SELECT * FROM users`// SQl query
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var users []models.User //return users slice
	for rows.Next() {
		var user models.User 
		// unmarshal the row object to user
		err := rows.Scan(&user.ID, &user.Email, &user.PhoneNumber, &user.FirstName, &user.LastName, &user.NickName, &user.Password)
		if err != nil {
			log.Printf("Unable to scan the row. %v", err)
        	continue // optionally skip the bad row instead of failing
		}

		users = append(users, user)
	}

	//return empty user on error 
	return users, err
}