package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/guiidc/todo-list/entities"
)

func TranslateValidationErrors(e validator.ValidationErrors) HttpError {
	errorsList := []entities.CustomError{}

	for _, value := range e {
		switch value.Tag() {
		case "required":
			errorsList = append(errorsList, entities.CustomError{
				Field:   value.Field(),
				Message: fmt.Sprintf("O campo %s é obrigatório", value.Field()),
				Value:   value.Value().(string),
			})
		case "min":
			errorsList = append(errorsList, entities.CustomError{
				Field:   value.Field(),
				Message: fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres", value.Field(), value.Param()),
				Value:   value.Value().(string),
			})
		case "max":
			errorsList = append(errorsList, entities.CustomError{
				Field:   value.Field(),
				Message: fmt.Sprintf("O campo %s deve ter no máximo %s caracteres", value.Field(), value.Param()),
				Value:   value.Value().(string),
			})
		case "email":
			errorsList = append(errorsList, entities.CustomError{
				Field:   value.Field(),
				Message: fmt.Sprintf("O campo %s deve ser um email válido", value.Field()),
				Value:   value.Value().(string),
			})
		case "eqfield":
			errorsList = append(errorsList, entities.CustomError{
				Field:   value.Field(),
				Message: fmt.Sprintf("O campo %s deve ser igual ao campo %s", value.Field(), value.Param()),
				Value:   value.Value().(string),
			})
		}
	}
	return HttpError{
		Errors: errorsList,
	}
}
