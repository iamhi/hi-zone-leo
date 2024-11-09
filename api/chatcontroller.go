package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhi/leo/internal/ollamahandler"
	ollamahander_requests "github.com/iamhi/leo/internal/ollamahandler/requests"
)

type sendMessageInput struct {
	Message string `json:"message"`
}

func initializeChatController(chat_group_router *gin.RouterGroup) {
	chat_group_router.POST("/", SendMessage)
}

func SendMessage(ctx *gin.Context) {
	var input sendMessageInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chatResponse, _ := ollamahandler.SendChatMessage(input.Message, []ollamahander_requests.MessageRequest{})

	ctx.JSON(http.StatusBadRequest, chatResponse.Message)

	return
}
