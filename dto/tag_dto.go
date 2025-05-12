package dto

type TagDto struct {
	Id      int64            `json:"id"`
	Name    string           `json:"name"`
	Color   int              `json:"color"`
	Owner   GenericEntityDto `json:"owner"`
	Project GenericEntityDto `json:"project"`
}
