package handlers

import (
	"encoding/json"
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

	// return the created item
	json.NewEncoder(w).Encode(item)
}

// -------------- Get item by ID --------------
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	
	item, err := models.GetItem(int64(id))
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	// send user with matching id
	json.NewEncoder(w).Encode(item)
}

// -------------- Update Item by ID --------------
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Get the item ID from URL params
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var itemData models.ItemData;
	
	if err := json.NewDecoder(r.Body).Decode(&itemData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	item, err := models.UpdateItem(
		id,
		itemData.Name,
		itemData.Description,
		itemData.Image,
		itemData.Price,
		itemData.Quantity,
		itemData.Available,
	)
	
	if err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(item)
}

// -------------- Get item by ID --------------
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	isDeleted, err := models.DeleteItem(int64(id))
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return 
	}

	if !isDeleted {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Item successfully deleted",
	})
}

// -------------- Search an Item --------------
func SearchItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Parse query parameters
    query := r.URL.Query().Get("query")
    
    var params models.SearchParams
    params.Query = query

    // Parse optional filters
    if categoryID := r.URL.Query().Get("category_id"); categoryID != "" {
        if catID, err := strconv.Atoi(categoryID); err == nil {
            params.CategoryID = &catID
        }
    }

    if minPrice := r.URL.Query().Get("min_price"); minPrice != "" {
        if price, err := strconv.Atoi(minPrice); err == nil {
            params.MinPrice = &price
        }
    }

    if maxPrice := r.URL.Query().Get("max_price"); maxPrice != "" {
        if price, err := strconv.Atoi(maxPrice); err == nil {
            params.MaxPrice = &price
        }
    }

    if available := r.URL.Query().Get("available"); available != "" {
        isAvailable := available == "true"
        params.Available = &isAvailable
    }

    // Perform search
    items, err := models.SearchItems(params)
    if err != nil {
        http.Error(w, "Failed to search items: "+err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(items)
}

