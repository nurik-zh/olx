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
	routes.RegisterCategoryRoutes(r)

	r.Run(":8080")
}
