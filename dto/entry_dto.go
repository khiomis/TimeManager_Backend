package dto

import (
	"github.com/google/uuid"
	"time"
)

type EntryDto struct {
	Id          uuid.UUID        `json:"id"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	StartAt     time.Time        `json:"startAt"`
	FinishAt    time.Time        `json:"finishAt"`
	Project     ProjectDto       `json:"project"`
	Task        TaskDto          `json:"task"`
	Owner       GenericEntityDto `json:"owner"`
}
