package handlers

import (
	"mongo-messages-api/models"
	"mongo-messages-api/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateMessageHandler(c *gin.Context) {
	var msg models.Message

	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	msg.CreatedAt = time.Now()

	err := repository.CreateMessage(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create message"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "message created"})

}

func GetMessagesByRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")

	messages, err := repository.GetMessagesByRoom(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
	}
	c.JSON(http.StatusOK, messages)
}

func UpdateMessageHandler(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Text string `json:"text"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	err := repository.UpdateMessage(id, body.Text)
	if err != nil {
		c.JSON(500, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteMessageHandler(c *gin.Context) {

	id := c.Param("id")

	err := repository.DeleteMessage(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "delete failed"})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}