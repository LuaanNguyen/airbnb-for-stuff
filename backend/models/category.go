package models

type Category struct {
    ID          int64  `json:"id" db:"c_id"`
    Name        string `json:"name" db:"c_name"`
    Description string `json:"description" db:"c_description"`
}
