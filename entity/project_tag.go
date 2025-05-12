package entity

type ProjectTag struct {
	IdProject int64 `json:"idProject" db:"id_project"`
	IdTag     int64 `json:"idTag" db:"id_tag"`
}
