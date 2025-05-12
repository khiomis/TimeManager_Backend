package dto

import (
	"backend_time_manager/entity"
	"github.com/google/uuid"
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
