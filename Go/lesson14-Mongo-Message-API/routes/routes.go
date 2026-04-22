package routes

import (
	"mongo-messages-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/messages", handlers.CreateMessageHandler)
	r.GET("/messages/:room_id", handlers.GetMessagesByRoomHandler)
	r.PUT("/messages/:id", handlers.UpdateMessageHandler)
	r.DELETE("/messages/:id", handlers.DeleteMessageHandler)
}
