package dto

import (
	"backend_time_manager/entity"
	"github.com/google/uuid"
	"time"
)

type SaveProjectDto struct {
	Name  string   `json:"name"`
	Color int      `json:"color"`
	Tags  []TagDto `json:"tags"`
}

type ProjectDto struct {
	Id        uuid.UUID        `json:"id"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Name      string           `json:"name"`
	Color     int              `json:"color"`
	Owner     GenericEntityDto `json:"owner"`
	Tags      []TagDto         `json:"tags"`
}

func (dto ProjectDto) From(project entity.Project) ProjectDto {
	dto.CreatedAt = project.CreatedAt
	dto.UpdatedAt = project.UpdatedAt
	dto.Name = project.Name
	dto.Color = project.Color
	return dto
}
