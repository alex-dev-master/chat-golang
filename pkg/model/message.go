package model

import "time"

type Message struct {
	Id         int    `json:"id" db:"id"`
	UserId     int    `json:"user_id" db:"user_id"`
	ChatRubric int    `json:"chat_rubric_id" db:"chat_rubric_id"`
	Content    string `json:"content" binding:"required"`
	Created    time.Time  `json:"created"`
	Disabled   bool `json:"disabled"`
}

type MessageWithUser struct {
	Id        int    `json:"id" db:"id"`
	UserId    int    `json:"user_id" db:"user_id"`
	Content    string `json:"content" binding:"required"`
	Created    time.Time  `json:"created"`
	ChatRubric int    `json:"chat_rubric_id" db:"chat_rubric_id"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"-"`
	Token     string `json:"-"`
	Disabled  string `json:"disabled"`
}
