package service

import "myproject/forum/server/repository"

type IReactionService interface {
}

type ReactionService struct {
	ReactionRepository repository.IReactionRepository
}

func NewReactionService(reactionRepository repository.IReactionRepository) *ReactionService {
	return &ReactionService{
		ReactionRepository: reactionRepository,
	}
}
