package chatcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhi/leo/api/middlewares"
	"github.com/iamhi/leo/internal/chathandler"
	"github.com/iamhi/leo/internal/userhandler"
)

type sendMessageRequest struct {
	Content string `json:"content"`
}

func getMessages(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if exists {
		user_details, ok := user_details_obj.(userhandler.UserDetails)

		if ok {
			fmt.Printf("%s\n", user_details.Uuid)
			chat_dto, chat_handler_error := chathandler.GetChat(user_details)

			if chat_handler_error != nil {
				context.JSON(http.StatusBadRequest,
					gin.H{"errorCode": chat_handler_error.GetCode(),
						"error": chat_handler_error.Error()})

				return
			}

			context.JSON(http.StatusOK, chat_dto)

			return
		}
	}

	context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get chat or user details"})
}

func sendMessage(context *gin.Context) {
	user_details_obj, exists := context.Get(middlewares.USER_DETAILS)

	if exists {
		user_details, ok := user_details_obj.(userhandler.UserDetails)

		if ok {
			var request_body sendMessageRequest

			if err := context.ShouldBindJSON(&request_body); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

				return
			}

			chat_dto, chat_handler_error := chathandler.SendMessage(user_details, request_body.Content)

			if chat_handler_error != nil {
				context.JSON(http.StatusBadRequest,
					gin.H{"errorCode": chat_handler_error.GetCode(),
						"error": chat_handler_error.Error()})

				return
			}

			context.JSON(http.StatusOK, chat_dto)

			return
		}
	}

	context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get chat or user details"})
}

const CHAT_CONTROLLER_PREFIX = "/chat"

func InitializeChatController(parent_router_group *gin.RouterGroup) {
	chat_router_group := parent_router_group.Group(CHAT_CONTROLLER_PREFIX)

	chat_router_group.GET("/", middlewares.Authorize(), getMessages)
	chat_router_group.POST("/", middlewares.Authorize(), sendMessage)
}
