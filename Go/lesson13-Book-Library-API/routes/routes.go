package routes

import (
	"lesson13-Book-Library-API/handlers"
	"lesson13-Book-Library-API/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/books", handlers.CreateBook)
		protected.PUT("/books/:id", handlers.UpdateBook)
		protected.DELETE("/books/:id", handlers.DeleteBook)
	}
}