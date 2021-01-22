package model

type ChatRubric struct {
	Id       int    `json:"-" db:"id"`
	UserId   int    `json:"user_id" db:"user_id"`
	Caption  string `json:"caption" binding:"required"`
	Disabled string `json:"disabled"`
}

type ChatRubricUser struct {
	Id        int    `json:"-" db:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	Caption   string `json:"caption" binding:"required"`
	Disabled  string `json:"disabled"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"-"`
	Token     string `json:"-"`
}
