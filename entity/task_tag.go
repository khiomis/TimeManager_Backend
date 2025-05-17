package entity

type TaskTag struct {
	IdTask int64 `json:"id_task" db:"id_task"`
	IdTag  int64 `json:"id_tag" db:"id_tag"`
}
