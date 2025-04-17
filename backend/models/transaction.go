package models

import "time"

type Transaction struct {
    ID        int64     `json:"id" db:"t_id"`
    UserID    int64     `json:"user_id" db:"u_id"`
    Type      string    `json:"type" db:"t_type"` // ENUM: 'Purchase', 'Sale', etc.
    ItemID    int64     `json:"item_id" db:"i_id"`
    Date      time.Time `json:"date" db:"t_date"`
    Amount    int       `json:"amount" db:"t_amount"`
}
