package main

import (
	"mongo-messages-api/config"
	"mongo-messages-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectMongo()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
