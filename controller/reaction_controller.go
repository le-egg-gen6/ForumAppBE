package controller

import (
	"myproject/forum/service"
)

type ReactionController struct {
	ReactionService service.IReactionService
}

func NewReactionController(reactionService service.IReactionService) *ReactionController {
	return &ReactionController{
		ReactionService: reactionService,
	}
}
