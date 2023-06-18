package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guiidc/todo-list/controllers"
	service "github.com/guiidc/todo-list/services"
)

var (
	authService    = service.NewAuthService()
	authController = controllers.NewAuthController(authService)
)

func MakeAuthRoutes(r *gin.Engine) {
	r.POST("/register", authController.Register)
}
