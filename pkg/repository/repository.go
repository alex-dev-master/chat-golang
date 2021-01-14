package repository

import (
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GetUser(username, password string) (model.User, error)

}

type ChatRubric interface {
	CreateRubric(rubric model.ChatRubric) (int64, error)
	GetRubrics() ([]model.ChatRubricUser, error)
}

type Repository struct {
	Authorization
	ChatRubric
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		ChatRubric: NewChatRubric(db),
	}
}
