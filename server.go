package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guiidc/todo-list/routes"
)

func main() {
	server := gin.Default()
	routes.MakeAuthRoutes(server)

	server.Run(":8080")
}
