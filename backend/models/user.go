package models

type User struct {
    ID          int64  `json:"id" db:"u_id"`
    Email       string `json:"email" db:"u_email"`
    PhoneNumber string `json:"phone_number" db:"u_phone_number"`
    FirstName   string `json:"first_name" db:"u_first_name"`
    LastName    string `json:"last_name" db:"u_last_name"`
    NickName    *string `json:"nick_name,omitempty" db:"u_nick_name"` // nullable
    Password    string `json:"-" db:"u_password"` // hide in JSON responses
}
