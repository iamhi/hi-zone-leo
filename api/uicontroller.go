package api

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/iamhi/leo/views"
)

type uiController struct {
}

func renderTempl(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func initializeUiController(ui_group_router *gin.RouterGroup) {
	ui_group_router.GET(CHAT_CONTROLLER_PREFIX, func(ctx *gin.Context) {
		renderTempl(ctx, http.StatusOK, views.Chat())
	})

	ui_group_router.GET("/", func(ctx *gin.Context) {
		renderTempl(ctx, http.StatusOK, views.Index())
	})

	ui_group_router.GET("/components/chat", func(ctx *gin.Context) {
		views.Chat()
		renderTempl(ctx, http.StatusOK, views.ChatContent())
	})
}
