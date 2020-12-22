package service

import (
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/alex-dev-master/chat-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ChatRubric interface {
	Create(rubric model.ChatRubric) (int64, error)
}

type Service struct {
	Authorization
	ChatRubric
}

func NewService(repos *repository.Repository) *Service  {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ChatRubric: NewChatRubricService(repos.ChatRubric),
	}
}