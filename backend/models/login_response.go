package models

type LoginResponse struct {
	Token     string `json:"token"`
	UserID    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
