package controller

import "myproject/forum/server/service"

type PostController struct {
	PostService service.IPostService
}

func NewPostController(postService service.IPostService) *PostController {
	return &PostController{}
}
