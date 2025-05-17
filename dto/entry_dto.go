package dto

import (
	"github.com/google/uuid"
	"time"
)

type EntryDto struct {
	Id          uuid.UUID        `json:"id"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	StartAt     time.Time        `json:"start_at"`
	FinishAt    time.Time        `json:"finish_at"`
	Project     ProjectDto       `json:"project"`
	Task        TaskDto          `json:"task"`
	Owner       GenericEntityDto `json:"owner"`
}
