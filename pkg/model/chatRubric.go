package model

type ChatRubric struct {
	Id        int    `json:"-" db:"id"`
	UserId        int    `json:"user_id" db:"user_id"`
	Caption  string `json:"caption" binding:"required"`
	Disabled  string `json:"disabled"`
}
