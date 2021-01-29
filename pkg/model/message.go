package model

import "time"

type Message struct {
	Id         int    `json:"-" db:"id"`
	UserId     int    `json:"user_id" db:"user_id"`
	ChatRubric int    `json:"chat_rubric_id" db:"chat_rubric_id"`
	Content    string `json:"content" binding:"required"`
	Created    time.Time  `json:"created"`
	Disabled   bool `json:"disabled"`
}

