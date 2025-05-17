package entity

type ProjectTag struct {
	IdProject int64 `json:"id_project" db:"id_project"`
	IdTag     int64 `json:"id_tag" db:"id_tag"`
}
