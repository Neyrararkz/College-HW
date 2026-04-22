package main

import (
	db "notes-api/config"
	"notes-api/handlers"
	"notes-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.POST("/notes", handlers.CreateNote)
	auth.GET("/my-notes", handlers.GetMyNotes)
	auth.GET("/notes/:id", handlers.GetNote)
	auth.PUT("/notes/:id", handlers.UpdateNote)
	auth.DELETE("/notes/:id", handlers.DeleteNote)

	r.Run(":8080")
}
