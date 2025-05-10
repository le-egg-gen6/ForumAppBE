package handler

import (
	"forum/models"
	"forum/repository"
	"forum/server/http_server/middlewares"
	"forum/shared"
	"forum/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitializeReactionHandler(router *gin.RouterGroup) {
	reactionRouter := router.Group("/reaction")
	{
		reactionRouter.GET("/post", middlewares.AuthenticationMiddlewares(), ReactionToPost)
		reactionRouter.GET("/comment", middlewares.AuthenticationMiddlewares(), ReactionToComment)
		reactionRouter.GET("/story", middlewares.AuthenticationMiddlewares(), ReactionToStory)
	}
}

func ReactionToPost(c *gin.Context) {
	userID := utils.GetCurrentContextUserID(c)
	if userID < 0 {
		shared.SendUnauthorized(c)
		return
	}

	user, err := repository.GetUserRepositoryInstance().FindByID(uint(userID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if user == nil {
		shared.SendUnauthorized(c)
		return
	}

	postIDStr := utils.GetRequestParam(c, "id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		shared.SendBadRequest(c, "Post not exist")
		return
	}
	post, err := repository.GetPostRepositoryInstance().FindByID(uint(postID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if post == nil {
		shared.SendBadRequest(c, "Post not exist")
		return
	}
	reactionType := utils.GetRequestParam(c, "type")
	if utils.IsReactionTypeValid(reactionType) {
		shared.SendBadRequest(c, "Reaction type not allowed")
		return
	}
	reaction := post.GetReaction(reactionType)
	if reaction == nil {
		reaction = &models.ContentReaction{
			Type:   reactionType,
			PostID: &post.ID,
			Count:  0,
		}
		reaction, err = repository.GetReactionRepositoryInstance().Create(reaction)
		if err != nil {
			shared.SendInternalServerError(c)
			return
		}
	}
	reaction.Count = reaction.Count + 1
	err = repository.GetReactionRepositoryInstance().Update(reaction)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	shared.SendSuccess(c, "Ok!")
}

func ReactionToComment(c *gin.Context) {
	userID := utils.GetCurrentContextUserID(c)
	if userID < 0 {
		shared.SendUnauthorized(c)
		return
	}

	user, err := repository.GetUserRepositoryInstance().FindByID(uint(userID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if user == nil {
		shared.SendUnauthorized(c)
		return
	}

	commentIDStr := utils.GetRequestParam(c, "id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		shared.SendBadRequest(c, "Post not exist")
		return
	}
	comment, err := repository.GetCommentRepositoryInstance().FindByID(uint(commentID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if comment == nil {
		shared.SendBadRequest(c, "Comment not exist")
		return
	}
	reactionType := utils.GetRequestParam(c, "type")
	if utils.IsReactionTypeValid(reactionType) {
		shared.SendBadRequest(c, "Reaction type not allowed")
		return
	}
	reaction := comment.GetReaction(reactionType)
	if reaction == nil {
		reaction = &models.ContentReaction{
			Type:      reactionType,
			CommentID: &comment.ID,
			Count:     0,
		}
		reaction, err = repository.GetReactionRepositoryInstance().Create(reaction)
		if err != nil {
			shared.SendInternalServerError(c)
			return
		}
	}
	reaction.Count = reaction.Count + 1
	err = repository.GetReactionRepositoryInstance().Update(reaction)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	shared.SendSuccess(c, "Ok!")
}

func ReactionToStory(c *gin.Context) {

}
