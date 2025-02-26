package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/dtos"
	"myproject/forum/service"
	"myproject/forum/shared"
	"myproject/forum/util"
	"net/http"
)

type PostController struct {
	PostService service.IPostService
}

func NewPostController(postService service.IPostService) *PostController {
	return &PostController{
		PostService: postService,
	}
}

func (pc *PostController) CreateNewPost(c *gin.Context) {
	var postDTO dtos.SimplePostDTO
	if err := c.ShouldBindJSON(&postDTO); err != nil {
		shared.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := util.GetCurrentContextUserID(c)
	if userID == -1 || userID != postDTO.Author.ID {
		shared.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	createdPost, err := pc.PostService.CreatePost(&postDTO)
	if err != nil {
		shared.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	shared.SendSuccess(c, createdPost)
}
