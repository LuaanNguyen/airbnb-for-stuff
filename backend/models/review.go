package models

type Review struct {
    ID      int64  `json:"id" db:"r_id"`
    Comment string `json:"comment" db:"r_comment"`
    Star    int    `json:"star" db:"r_star"`
    UserID  int64  `json:"user_id" db:"u_id"`
}
