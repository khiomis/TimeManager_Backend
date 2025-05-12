package entity

import (
	"github.com/google/uuid"
	"time"
)

type Entry struct {
	Id          int64     `json:"id" db:"id_entry"`
	Uuid        uuid.UUID `json:"uuid" db:"id_entry_uuid"`
	CreatedAt   time.Time `json:"createdAt" db:"dt_created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"dt_updated_at"`
	Name        string    `json:"name" db:"nm_entry"`
	Description string    `json:"description" db:"ds_description"`
	StartAt     time.Time `json:"startAt" db:"dt_start_at"`
	FinishAt    time.Time `json:"finishAt" db:"dt_finish_at"`
	IdOwner     int64     `json:"idOwner" db:"id_owner"`
	IdProject   int64     `json:"idProject" db:"id_project"`
	IdTask      int64     `json:"idTask" db:"id_task"`
}
