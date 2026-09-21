package routes

import (
	"notes-api/handlers"
	"notes-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/chats", handlers.CreateChat)
		auth.GET("/chats", handlers.GetMyChats)

		auth.POST("/messages", handlers.SendMessage)
		auth.GET("/messages/:chat_id", handlers.GetMessages)
	}
}