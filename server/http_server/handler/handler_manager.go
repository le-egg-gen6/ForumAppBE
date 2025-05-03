package handler

import "github.com/gin-gonic/gin"

func InitializeHandler(routerGroup *gin.RouterGroup) {
	InitializeAuthHandler(routerGroup)
	InitializePostHandler(routerGroup)
	InitializeCommentHandler(routerGroup)
	InitializeReactionHandler(routerGroup)
}
