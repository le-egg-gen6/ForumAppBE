package shared

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: "Success",
		Data:    data,
	})
}

func SendError(c *gin.Context, code int, msg string) {
	c.JSON(code, APIResponse{
		Status:  "error",
		Message: msg,
		Data:    nil,
	})
}

func SendInternalServerError(c *gin.Context) {
	SendError(c, http.StatusInternalServerError, "An error occurred, please try again later")
}

func SendUnauthorized(c *gin.Context) {
	SendError(c, http.StatusUnauthorized, "Unauthorized")
}

func SendBadRequest(c *gin.Context, msg string) {
	SendError(c, http.StatusBadRequest, msg)
}
