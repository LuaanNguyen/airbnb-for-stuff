package models

import (
	"fmt"

	"github.com/LuaanNguyen/backend/db"
)

// -------------- GetAllUsers retrieves all users from the database --------------
func GetAllUsers() ([]User, error) {
	rows, err := db.DB.Query("SELECT u_id, u_email, u_phone_number, u_first_name, u_last_name, u_nick_name FROM users")
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Email, &u.PhoneNumber, &u.FirstName, &u.LastName, &u.NickName)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, u)
	}

	return users, nil
}

// -------------- GetUser retrieves a single user by ID --------------
func GetUser(id int64) (User, error) {
	var u User
	err := db.DB.QueryRow("SELECT u_id, u_email, u_phone_number, u_first_name, u_last_name, u_nick_name FROM users WHERE u_id = $1", id).
		Scan(&u.ID, &u.Email, &u.PhoneNumber, &u.FirstName, &u.LastName, &u.NickName)
	if err != nil {
		return User{}, fmt.Errorf("error querying user: %v", err)
	}
	return u, nil
}

// -------------- GetAllItems retrieves all items from the database --------------
func GetAllItems() ([]Item, error) {
	rows, err := db.DB.Query("SELECT i_id, i_name, i_description, i_image, c_id, i_price, i_date_listed, i_quantity, i_available FROM items")
	if err != nil {
		return nil, fmt.Errorf("error querying items: %v", err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var i Item
		err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Image, &i.CategoryID, &i.Price, &i.DateListed, &i.Quantity, &i.Available)
		if err != nil {
			return nil, fmt.Errorf("error scanning item: %v", err)
		}
		items = append(items, i)
	}

	return items, nil
}

// -------------- GetUserByEmail retrieves a user email for login --------------
func GetUserByEmail(email string) (User, error) {
	var user User 
	err := db.DB.QueryRow(`
		SELECT u_id, u_email, u_phone_number, u_first_name, u_last_name, u_nick_name, u_password 
        FROM users 
        WHERE u_email = $1`, email).
		Scan(&user.ID, &user.Email, &user.PhoneNumber, &user.FirstName, &user.LastName, &user.NickName, &user.Password)
	if err != nil {
		return User{}, fmt.Errorf("error querying user: %v", err)
	}

	return user, nil
}


// -------------- Get all rental items that are available for rent (Limit to 50 for now) --------------
func GetAvailableItemsWithOwners() ([]ItemWithOwner, error) {
    query := `
        SELECT 
            i.i_id, 
            i.i_name, 
            i.i_description, 
            i.i_price,
            i.owner_id,
            CONCAT(u.u_first_name, ' ', u.u_last_name) as owner_name,
            i.i_available
        FROM items i
        JOIN users u ON i.owner_id = u.u_id
        WHERE i.i_available = true
        AND NOT EXISTS (
            SELECT 1 FROM rentals r
            WHERE r.i_id = i.i_id 
            AND r.status = 'approved'
            AND r.end_date > CURRENT_TIMESTAMP
        ) 
		LIMIT 50`

    rows, err := db.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error querying available items: %v", err)
    }
    defer rows.Close()

    var items []ItemWithOwner
    for rows.Next() {
        var item ItemWithOwner
        err := rows.Scan(
            &item.ID, 
            &item.Name, 
            &item.Description, 
            &item.Price, 
            &item.OwnerID, 
            &item.OwnerName, 
            &item.Available,
        )
        if err != nil {
            return nil, fmt.Errorf("error scanning item: %v", err)
        }
        items = append(items, item)
    }
    return items, nil
}

// -------------- Create a rental request --------------
func CreateRentalRequest(rental *RentalRequest) error {
    query := `
        INSERT INTO rentals (item_id, renter_id, start_date, end_date, status, total_price)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING rental_id`
    
    err := db.DB.QueryRow(
        query,
        rental.ItemID,
        rental.RenterID,
        rental.StartDate,
        rental.EndDate,
        "pending",
        rental.TotalPrice,
    ).Scan(&rental.ID)
    
    return err
}


// -------------- Get all category--------------
func GetAllCategories() ([]Category, error) {
    var categories []Category
    rows, err := db.DB.Query(`SELECT * FROM categories`)
    if err != nil {
        return nil, fmt.Errorf("error querying categories: %v", err )
    }

    for rows.Next() {
		var m Category
		err := rows.Scan(&m.ID, &m.Name, &m.Description)
		if err != nil {
			return nil, fmt.Errorf("error scanning category: %v", err)
		}
		categories = append(categories, m)
	}

	return categories, nil
}


// -------------- Create a new item --------------
func CreateItem(item *Item) error {
    query := `
        INSERT INTO items (i_name, i_description, i_image, c_id, owner_id, i_price, i_date_listed, i_quantity, i_available)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING i_id`  // This will return the auto-generated ID

    // Notice i_id is NOT in the field list above

    err := db.DB.QueryRow(
        query,
        item.Name,
        item.Description,
        item.Image,
        item.CategoryID,
        item.OwnerID,
        item.Price,
        item.DateListed,
        item.Quantity,
        item.Available,
    ).Scan(&item.ID)
    
    if err != nil {
        return fmt.Errorf("error creating item: %v", err)
    }
    
    return nil
}

// -------------- GetItem retrieves a single item by ID --------------
func GetItem(id int64) (Item, error) {
	var i Item
	err := db.DB.QueryRow("SELECT i_id, i_name, i_description, i_price, i_date_listed, i_quantity, i_available FROM items WHERE i_id = $1", id).
		Scan(&i.ID, &i.Name, &i.Description, &i.Price, &i.DateListed, &i.Quantity, &i.Available)
	if err != nil {
		return Item{}, fmt.Errorf("error querying user: %v", err)
	}
	return i, nil
}

// -------------- Delete an item by its ID --------------
func DeleteItem(id int64) (bool, error) {
    result, err := db.DB.Exec("DELETE FROM items WHERE i_id = $1", id)
    if err != nil {
        return false, err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false, err
    }
    
    return rowsAffected > 0, nil
}