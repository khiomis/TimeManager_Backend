package dto

import (
	"github.com/google/uuid"
	"time"
)

type ProjectDto struct {
	Id        uuid.UUID        `json:"id"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Name      string           `json:"name"`
	Color     int              `json:"color"`
	Owner     GenericEntityDto `json:"owner"`
	Tags      []TagDto         `json:"tags"`
}
