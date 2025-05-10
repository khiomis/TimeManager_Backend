package dto

import (
	"backend_time_manager/entity"
	"time"
)

type UserDTO struct {
	Id        int64             `json:"id"`
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
