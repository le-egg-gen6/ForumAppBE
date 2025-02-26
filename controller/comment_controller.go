package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/dtos"
	"myproject/forum/service"
	"myproject/forum/shared"
	"myproject/forum/util"
	"net/http"
)

type CommentController struct {
	CommentService *service.CommentService
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

func (cc *CommentController) CreateNewComment(c *gin.Context) {
	var commentDTO dtos.SimpleCommentDTO
	if err := c.ShouldBindJSON(&commentDTO); err != nil {
		shared.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := util.GetCurrentContextUserID(c)
	if userID == -1 || userID != commentDTO.Author.ID {
		shared.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	createdComment, err := cc.CommentService.CreateComment(&commentDTO)
	if err != nil {
		shared.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	shared.SendSuccess(c, createdComment)

}
