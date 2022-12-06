package service

import (
	"crypto_app/pkg/models"
	"crypto_app/pkg/repository"
)

type RepliesService struct {
	repository repository.Replies
}

func NewRepliesService(repository repository.Replies) *RepliesService {
	return &RepliesService{repository: repository}
}

func (r *RepliesService) CreateReply(reply models.Reply) (int, error) {
	return r.repository.CreateReply(reply)
}
