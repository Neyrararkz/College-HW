package main

import(
	"notes-api/config"
	"notes-api/routes"

	"github.com/gin-gonic/gin"
)
func main() {
	config.ConnectDB()
	config.ConnectMongo()

	r := gin.Default()

	routes.SetupRoutes(r)
	r.Run(":8080")
}