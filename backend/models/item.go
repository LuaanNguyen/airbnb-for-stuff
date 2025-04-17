package models

import "time"

type Item struct {
    ID          int64      `json:"id" db:"i_id"`
    Name        string     `json:"name" db:"i_name"`
    Description string     `json:"description" db:"i_description"`
    Image       *[]byte    `json:"image,omitempty" db:"i_image"` // nullable
    CategoryID  int64      `json:"category_id" db:"c_id"`
    Price       int        `json:"price" db:"i_price"`
    DateListed  time.Time  `json:"date_listed" db:"i_date_listed"`
    Quantity    int        `json:"quantity" db:"i_quantity"`
    Available   bool       `json:"available" db:"i_available"`
}
