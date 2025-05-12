package entity

type TaskTag struct {
	IdTask int64 `json:"idTask" db:"id_task"`
	IdTag  int64 `json:"idTag" db:"id_tag"`
}
