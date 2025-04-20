package models

import "time"

type RentalRequest struct {
    ID          int64     `json:"id"`
    ItemID      int64     `json:"item_id"`
    RenterID    int64     `json:"renter_id"`
    StartDate   time.Time `json:"start_date"`
    EndDate     time.Time `json:"end_date"`
    Status      string    `json:"status"`
    TotalPrice  int64     `json:"total_price"`
}