package service

import (
	"github.com/alex-dev-master/chat-golang/pkg/model"
	"github.com/alex-dev-master/chat-golang/pkg/repository"
)

type ChatRubricService struct {
	repo repository.ChatRubric
}

func NewChatRubricService(repo repository.ChatRubric) *ChatRubricService {
	return &ChatRubricService{repo: repo}
}

func (s *ChatRubricService) Create(rubric model.ChatRubric) (int64, error) {
	return s.repo.CreateRubric(rubric)
}
