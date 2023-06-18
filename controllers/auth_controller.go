package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/guiidc/todo-list/entities"
	"github.com/guiidc/todo-list/services"
	"github.com/guiidc/todo-list/utils"
)

type AuthController interface {
	Register(ctx *gin.Context)
}

type authController struct {
	AuthService services.AuthService
}

func NewAuthController(as services.AuthService) AuthController {
	return &authController{
		AuthService: as,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var newUser entities.CreateUserRequest

	if err := ctx.BindJSON(&newUser); err != nil {
		var ve validator.ValidationErrors

		if ok := errors.As(err, &ve); ok {
			translatedErrors := utils.TranslateValidationErrors(ve)
			ctx.JSON(400, translatedErrors)
			return
		}

		ctx.JSON(500, utils.InternalServerError())
		return
	}

	c.AuthService.Register(newUser)

	ctx.JSON(200, true)
}
