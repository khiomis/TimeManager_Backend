package entity

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id        int64     `json:"id" db:"id_task"`
	Uuid      uuid.UUID `json:"uuid" db:"id_task_uuid"`
	CreatedAt time.Time `json:"created_at" db:"dt_created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"dt_updated_at"`
	Name      string    `json:"name" db:"nm_task"`
	IdProject int64     `json:"id_project" db:"id_project"`
	IdOwner   int64     `json:"id_owner" db:"id_owner"`
}
