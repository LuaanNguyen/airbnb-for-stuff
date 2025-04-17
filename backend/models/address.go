package models

type Address struct {
    ID       int64  `json:"id" db:"a_id"`
    UserID   int64  `json:"user_id" db:"u_id"`
    Street   string `json:"street" db:"a_street"`
    City     string `json:"city" db:"a_city"`
    State    string `json:"state" db:"a_state"`
    Zipcode  string `json:"zipcode" db:"a_zipcode"`
    Country  string `json:"country" db:"a_country"`
}
