package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/LuaanNguyen/backend/middleware"
	"github.com/LuaanNguyen/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// -------------- Check health --------------
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Message: "Hello, you have successfully connected to Postgres 🫶",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// -------------- Get all users --------------
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get all the users in the db 
	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve all users", http.StatusInternalServerError)
		return
	}

	// send all users as response
	json.NewEncoder(w).Encode(users)
}

// -------------- Get user by ID --------------
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	user, err := models.GetUser(int64(id))
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	// send user with matching id
	json.NewEncoder(w).Encode(user)
}

// -------------- Get all items --------------
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, err := models.GetAllItems()
	if err != nil {
		http.Error(w, "Failed to retrieve all items", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

// -------------- login --------------
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Parse request body 
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	// Get user from database (use models function)
	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "Invalid Email", http.StatusUnauthorized)
		return 
	}

	// Verify password with bcypt (To be implemented)
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
    //     http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
    //     return
    // }

	if user.Password != req.Password {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
    	return
	}

	//Generate a JWT token 
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &models.Claims{
        UserID: user.ID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtKey := []byte(os.Getenv("JWT_SECRET"))
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        http.Error(w, "Error creating token", http.StatusInternalServerError)
        return
    }

    // Send response
    response := models.LoginResponse{
        Token:     tokenString,
        UserID:    user.ID,
        FirstName: user.FirstName,
        LastName:  user.LastName,
    }

    json.NewEncoder(w).Encode(response)
}

// -------------- Get avaialble items for rent --------------
func GetAvailableItems(w http.ResponseWriter, r *http.Request) {
    items, err := models.GetAvailableItemsWithOwners()
    if err != nil {
        http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(items)
}


// -------------- Get new rental request --------------
func CreateRentalRequest(w http.ResponseWriter, r *http.Request) {
    var req models.RentalRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    // Get user ID from JWT token
    userID, _ := middleware.GetUserIDFromContext(r)
    req.RenterID = int64(userID)
    
    if err := models.CreateRentalRequest(&req); err != nil {
        http.Error(w, "Failed to create rental request", http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(req)
}

// -------------- Get all categories --------------
func GetAllCategories(w http. ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := models.GetAllCategories()
	if err != nil {
		http.Error(w, "Failed to retrieve all users", http.StatusInternalServerError)
		return
	}
	// send all users as response
	json.NewEncoder(w).Encode(categories)
	
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update user handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// -------------- Create new rental items --------------
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// parse request body 
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return 
	}

	// get the user ID from JWT token and set as owner
	userID, _ := middleware.GetUserIDFromContext(r)
	item.OwnerID = int64(userID)


	// Set defaults 
	if item.DateListed.IsZero() {
		item.DateListed = time.Now()
	}

	// Create the item 
	if err := models.CreateItem(&item); err != nil {
		http.Error(w, "Failed to create item: "+err.Error(), http.StatusInternalServerError)
		return 
	}

	fmt.Println(item)
	// return the created item
	json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get item handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update item handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete item handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func SearchItems(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement search items handler
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}


