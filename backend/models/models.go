package models

import (
	"fmt"

	"github.com/LuaanNguyen/backend/db"
)

// GetAllUsers retrieves all users from the database
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

// GetUser retrieves a single user by ID
func GetUser(id int64) (User, error) {
	var u User
	err := db.DB.QueryRow("SELECT u_id, u_email, u_phone_number, u_first_name, u_last_name, u_nick_name FROM users WHERE u_id = $1", id).
		Scan(&u.ID, &u.Email, &u.PhoneNumber, &u.FirstName, &u.LastName, &u.NickName)
	if err != nil {
		return User{}, fmt.Errorf("error querying user: %v", err)
	}
	return u, nil
}

// GetAllItems retrieves all items from the database
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

// GetUserByEmail retrieves a user email for login 
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