package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iamhi/leo/api/controllers/chatcontroller"
	"github.com/iamhi/leo/api/controllers/usercontroller"
)

const SERVICE_PREFIX = "/hi-zone-api/leo"

const CHAT_CONTROLLER_PREFIX = "/chat"

func initialize(gin_engine *gin.Engine) {
	root_group := gin_engine.Group(SERVICE_PREFIX)

	root_group.Static("/static", "./static")

	// chat_group := root_group.Group(CHAT_CONTROLLER_PREFIX, middlewares.Authorize())

	// initializeChatController(chat_group)

	usercontroller.InitializeUserController(root_group)
	chatcontroller.InitializeChatController(root_group)

	ui_group := root_group.Group("/ui")

	initializeUiController(ui_group)
}
