package entity

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	Id        int64     `json:"id" db:"id_project"`
	Uuid      uuid.UUID `json:"uuid" db:"id_project_uuid"`
	CreatedAt time.Time `json:"createdAt" db:"dt_created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"dt_updated_at"`
	Name      string    `json:"name" db:"nm_project"`
	Color     int       `json:"color" db:"vl_color"`
	IdOwner   int64     `json:"idOwner" db:"id_owner"`
}
