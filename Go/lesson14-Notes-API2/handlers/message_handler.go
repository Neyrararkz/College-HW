package handlers

import (
	"net/http"
	"notes-api/models"
	"notes-api/mongo_repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(c *gin.Context) {
	var req CreateMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	chatID, _ := primitive.ObjectIDFromHex(req.ChatID)

	msg := models.Message{
		ChatID: chatID,
		UserID: c.GetInt("user_id"),
		Text:   req.Text,
	}

	err := mongo_repository.CreateMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "sent"})
}

func GetMessages(c *gin.Context) {
	chatIDParam := c.Param("chat_id")
	chatID, _ := primitive.ObjectIDFromHex(chatIDParam)

	messages, err := mongo_repository.GetMessages(chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}