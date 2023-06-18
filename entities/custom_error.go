package entities

type CustomError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value"`
}

type ErrorsResponse struct {
	Errors []CustomError `json:"errors"`
}
