package dto

import (
	"github.com/google/uuid"
	"time"
)

type TaskDto struct {
	Id        uuid.UUID        `json:"id"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Name      string           `json:"name"`
	Project   ProjectDto       `json:"project"`
	Owner     GenericEntityDto `json:"owner"`
	Tags      []TagDto         `json:"tags"`
}
