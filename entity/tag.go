package entity

import "time"

type Tag struct {
	Id        int64     `json:"id" db:"id_tag"`
	CreatedAt time.Time `json:"createdAt" db:"dt_created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"dt_updated_at"`
	Name      string    `json:"name" db:"nm_tag"`
	Color     int       `json:"color" db:"vl_color"`
	IdOwner   int64     `json:"idOwner" db:"id_owner"`
	IdProject int64     `json:"idProject" db:"id_project"`
}
