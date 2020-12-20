package service

import (
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/alex-dev-master/chat-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int64, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service  {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}