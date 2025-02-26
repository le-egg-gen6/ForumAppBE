package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/dtos"
	"myproject/forum/service"
	"myproject/forum/shared"
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

}
