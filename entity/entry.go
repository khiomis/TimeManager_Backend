package entity

import (
	"github.com/google/uuid"
	"time"
)

type Entry struct {
	Id          int64     `json:"id" db:"id_entry"`
	Uuid        uuid.UUID `json:"uuid" db:"id_entry_uuid"`
	CreatedAt   time.Time `json:"created_at" db:"dt_created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"dt_updated_at"`
	Name        string    `json:"name" db:"nm_entry"`
	Description string    `json:"description" db:"ds_description"`
	StartAt     time.Time `json:"start_at" db:"dt_start_at"`
	FinishAt    time.Time `json:"finish_at" db:"dt_finish_at"`
	IdOwner     int64     `json:"id_owner" db:"id_owner"`
	IdProject   int64     `json:"id_project" db:"id_project"`
	IdTask      int64     `json:"id_task" db:"id_task"`
}
