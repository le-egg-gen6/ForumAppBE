package handler

import (
	"forum/3rd_party_service/cloudinary_service"
	"forum/constant"
	"forum/dtos"
	"forum/models"
	"forum/repository"
	"forum/server/http_server/middlewares"
	"forum/shared"
	"forum/utils"
	"github.com/gin-gonic/gin"
)

func InitializeCommentHandler(router *gin.RouterGroup) {
	commentGroup := router.Group("/comment")
	{
		commentGroup.POST("/create", middlewares.AuthenticationMiddlewares(), CreateNewComment)
	}
}

func CreateNewComment(c *gin.Context) {
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

	var createCommentDTO dtos.SimpleCommentDTO
	if err := c.ShouldBind(&createCommentDTO); err != nil {
		shared.SendBadRequest(c, "Bad Request")
		return
	}

	post, err := repository.GetPostRepositoryInstance().FindByID(createCommentDTO.PostID)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if post == nil {
		shared.SendBadRequest(c, "Post Not Found")
		return
	}

	file, fileHeader, err := utils.GetFirstMultipartFile(c, constant.FileFormKey)
	if err != nil {
		shared.SendBadRequest(c, "Bad Request")
		return
	}
	if fileHeader != nil && !utils.IsFileValid(fileHeader) {
		shared.SendBadRequest(c, "Invalid file")
		return
	}

	comment := &models.Comment{
		UserID: &user.ID,
		PostID: &post.ID,
		Body:   createCommentDTO.Body,
	}
	if file != nil {
		fileDto := dtos.File{
			FileHeader: fileHeader,
			File:       file,
		}
		image, _ := cloudinary_service.GetFileUploaderInstance().UploadFile(&fileDto)
		if image != nil {
			comment.Image = image
		}
	}
	comment, err = repository.GetCommentRepositoryInstance().Create(comment)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	shared.SendSuccess(c, comment)
}
