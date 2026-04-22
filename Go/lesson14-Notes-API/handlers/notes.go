package handlers

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"notes-api/config"
)

func CreateNote(c *gin.Context) {

	var input struct {
		Title    string
		Content  string
		IsPublic bool
	}

	c.BindJSON(&input)

	userID := c.GetInt("user_id")

	_, err := db.DB.Exec(
		"INSERT INTO notes (title, content, is_public, user_id) VALUES ($1,$2,$3,$4)",
		input.Title, input.Content, input.IsPublic, userID,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "failed"})
		return
	}

	c.JSON(200, gin.H{"message": "created"})
}

func GetMyNotes(c *gin.Context) {

	userID := c.GetInt("user_id")

	rows, _ := db.DB.Query(
		"SELECT id, title, content, is_public FROM notes WHERE user_id=$1",
		userID,
	)
	defer rows.Close()

	var notes []gin.H

	for rows.Next() {
		var id int
		var title, content string
		var isPublic bool

		rows.Scan(&id, &title, &content, &isPublic)

		notes = append(notes, gin.H{
			"id": id,
			"title": title,
			"content": content,
			"is_public": isPublic,
		})
	}

	c.JSON(200, notes)
}

func GetNote(c *gin.Context) {

	id := c.Param("id")
	userID := c.GetInt("user_id")

	var title, content string
	var isPublic bool
	var ownerID int

	err := db.DB.QueryRow(
		"SELECT title, content, is_public, user_id FROM notes WHERE id=$1",
		id,
	).Scan(&title, &content, &isPublic, &ownerID)

	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	if !isPublic && ownerID != userID {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	c.JSON(200, gin.H{
		"title": title,
		"content": content,
	})
}

func UpdateNote(c *gin.Context) {

	id := c.Param("id")
	userID := c.GetInt("user_id")

	var input struct {
		Title    string
		Content  string
		IsPublic bool
	}

	c.BindJSON(&input)

	// проверяем владельца
	var ownerID int

	err := db.DB.QueryRow(
		"SELECT user_id FROM notes WHERE id=$1",
		id,
	).Scan(&ownerID)

	if err != nil {
		c.JSON(404, gin.H{"error": "note not found"})
		return
	}

	if ownerID != userID {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	// обновление
	_, err = db.DB.Exec(
		"UPDATE notes SET title=$1, content=$2, is_public=$3 WHERE id=$4",
		input.Title, input.Content, input.IsPublic, id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteNote(c *gin.Context) {

	id := c.Param("id")
	userID := c.GetInt("user_id")

	var ownerID int

	err := db.DB.QueryRow(
		"SELECT user_id FROM notes WHERE id=$1",
		id,
	).Scan(&ownerID)

	if err != nil {
		c.JSON(404, gin.H{"error": "note not found"})
		return
	}

	if ownerID != userID {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	_, err = db.DB.Exec(
		"DELETE FROM notes WHERE id=$1",
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "delete failed"})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}