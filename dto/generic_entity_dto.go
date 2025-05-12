package dto

import "github.com/google/uuid"

type GenericEntityDto struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
