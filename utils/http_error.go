package utils

import "github.com/guiidc/todo-list/entities"

type HttpError struct {
	Errors []entities.CustomError `json:"errors"`
}

func InternalServerError() HttpError {
	return HttpError{
		Errors: []entities.CustomError{
			{
				Field:   "server",
				Message: "Ocorreu um erro interno no servidor",
				Value:   "",
			},
		},
	}
}
