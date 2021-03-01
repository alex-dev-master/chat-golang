package handler

import (
	"github.com/alex-dev-master/chat-golang/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refreshToken)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/test", func(context *gin.Context) {
			context.JSON(http.StatusOK, map[string]interface{}{
				"test": "norm",
			})
		})

		rubrics := api.Group("/chat-rubrics")
		{
			rubrics.POST("/", h.createChatRubric)
			rubrics.GET("/", h.getChatRubrics)
		}

		messages := api.Group("/messages")
		{
			messages.POST("/:rubric", h.createMessage)
			messages.GET("/:rubric", h.getMessages)
		}

	}

	return router
}