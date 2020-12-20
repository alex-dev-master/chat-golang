package model

type User struct {
	Id        int    `json:"-" db:"id"`
	LastName  string `json:"last_name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Token  string `json:"token" binding:"required"`
	Disabled  string `json:"disabled" binding:"required"`
}
