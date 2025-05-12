package dto

import "github.com/google/uuid"

type TokenValidationDto struct {
	Email string    `json:"email"`
	Id    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}
