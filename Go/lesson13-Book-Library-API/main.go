package main

import (
	"lesson13-Book-Library-API/config"
	"lesson13-Book-Library-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}