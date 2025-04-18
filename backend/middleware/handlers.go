package middleware

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LuaanNguyen/backend/models"
)

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


func getUser(id int64) (models.User, error) {
	// create db connection 
	db :=CreateConnection()
	defer db.Close() 

	sqlStatement := `SELECT * FROM users WHERE u_id=$1`
	row := db.QueryRow(sqlStatement, id) //query

	var user models.User
	//unmarshal the row object to user 
	err := row.Scan(&user.ID, &user.Email, &user.PhoneNumber, &user.FirstName, &user.LastName, &user.NickName, &user.Password)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil 
	case nil:
		return user, nil 
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	// all success, return user without error
	return user, nil
}

func getAllItems() ([]models.Item, error) {
	db := CreateConnection()
	defer db.Close() 

	sqlStatement := `SELECT * FROM items`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.ID, 
			&item.Name, 
			&item.Description,
			&item.Image, 
			&item.CategoryID,
			&item.Price, 
			&item.DateListed,
			&item.Quantity,
			&item.Available,
		)
		if err != nil {
			log.Printf("Unable to scan the row. %v", err)
        	continue
		}
		items = append(items, item)
	}
	return items, nil
}