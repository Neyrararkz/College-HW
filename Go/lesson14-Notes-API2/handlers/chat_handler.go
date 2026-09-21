package handlers

import (
	"net/http"
	"notes-api/mongo_repository"

	"github.com/gin-gonic/gin"
)

type CreateChatRequest struct {
	UserIDs []int `json:"user_ids"`
}

func CreateChat(c *gin.Context) {
	var req CreateChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	chat, err := mongo_repository.CreateChat(req.UserIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create chat"})
		return
	}

	c.JSON(http.StatusOK, chat)
}

func GetMyChats(c *gin.Context) {
	userID := c.GetInt("user_id")

	chats, err := mongo_repository.GetUserChats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get chats"})
		return
	}

	c.JSON(http.StatusOK, chats)
}