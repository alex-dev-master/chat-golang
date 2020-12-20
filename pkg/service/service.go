package service

import "github.com/alex-dev-master/chat-golang/pkg/repository"

type Service struct {

}

func NewService(repos *repository.Repository) *Service  {
	return &Service{}
}