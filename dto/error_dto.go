package dto

type ErrorDto struct {
	Code        string          `json:"code"`
	Message     string          `json:"message"`
	Description string          `json:"description"`
	Errors      []FieldErrorDto `json:"errors"`
}

type FieldErrorDto struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (error ErrorDto) IsEmpty() bool {
	return len(error.Code) == 0 && len(error.Message) == 0 && len(error.Description) == 0 && len(error.Errors) == 0
}
