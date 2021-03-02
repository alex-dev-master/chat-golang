package service

import (
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/alex-dev-master/chat-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
	GetUser(username, password string) (model.User, error)
	GenerateToken(user model.User) (string, error)
	CreateRefreshToken(user model.User) (string, error)
	RefreshAccessToken(refreshToken string) (string, error)
	ParseToken(token string) (int, error)
}

type ChatRubric interface {
	Create(rubric model.ChatRubric) (int64, error)
	GetAll() ([]model.ChatRubricUser, error)
}

type Message interface {
	Create(message model.Message) (int, error)
	GetAllOfRubric(rubric int) ([]model.MessageWithUser, error)
}

type Service struct {
	Authorization
	ChatRubric
	Message
}

func NewService(repos *repository.Repository) *Service  {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ChatRubric: NewChatRubricService(repos.ChatRubric),
		Message: NewMessageService(repos.Message),
	}
}