package model

type User struct {
	Id           int    `json:"-" db:"id"`
	LastName     string `json:"last_name" binding:"required" db:"last_name"`
	FirstName    string `json:"first_name" binding:"required" db:"first_name"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Token        string `json:"token" db:"token"`
	ExpiredToken int64 `json:"expired_token" db:"expired_token"`
	Disabled     string `json:"disabled"`
}
