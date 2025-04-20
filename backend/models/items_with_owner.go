package models

type ItemWithOwner struct {
    ID          int64   `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price_per_day"`
    OwnerID     int64   `json:"owner_id"`
    OwnerName   string  `json:"owner_name"`
    Available   bool    `json:"available"`
}
