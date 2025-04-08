package main

import (
	"github.com/gin-gonic/gin"
	"olxkz/config"
	"olxkz/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	routes.RegisterProductRoutes(r)

	r.Run(":8080")
}
