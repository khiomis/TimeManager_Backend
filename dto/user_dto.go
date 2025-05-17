package dto

import (
	"backend_time_manager/constants"
	"backend_time_manager/entity"
	"backend_time_manager/utils"
	"github.com/google/uuid"
	"net/mail"
	"time"
)

type UserDTO struct {
	Id        uuid.UUID         `json:"id"`
	Email     string            `json:"email"`
	Name      string            `json:"name"`
	Status    entity.UserStatus `json:"status"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type CreateUserDto struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (user UserDTO) From(entity entity.User) UserDTO {
	user.Id = entity.Uuid
	user.Email = entity.Email
	user.Name = entity.Name
	user.Status = entity.Status
	user.UpdatedAt = entity.UpdatedAt
	return user
}

func (value CreateUserDto) Validate() ErrorDto {
	var errorDto = ErrorDto{}

	if len(value.Email) == 0 {
		errorDto.Errors = append(errorDto.Errors, FieldErrorDto{
			Field:   "email",
			Code:    constants.FormValueEmpty,
			Message: "Email is required",
		})
	} else if _, err := mail.ParseAddress(value.Email); err != nil {
		errorDto.Errors = append(errorDto.Errors, FieldErrorDto{
			Field:   "email",
			Code:    constants.FormValueInvalid,
			Message: "The email address is invalid",
		})
	}
	if len(value.Password) == 0 {
		errorDto.Errors = append(errorDto.Errors, FieldErrorDto{
			Field:   "password",
			Code:    constants.FormValueEmpty,
			Message: "Password is required",
		})
	} else if valid := utils.ValidatePassword(value.Password); !valid {
		errorDto.Errors = append(errorDto.Errors, FieldErrorDto{
			Field:   "password",
			Code:    constants.FormValueInvalid,
			Message: "The password is invalid",
		})
	}
	if len(value.Name) == 0 {
		errorDto.Errors = append(errorDto.Errors, FieldErrorDto{
			Field:   "name",
			Code:    constants.FormValueEmpty,
			Message: "Name is required",
		})
	}
	return errorDto
}
