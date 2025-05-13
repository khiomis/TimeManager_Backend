package database

import (
	"backend_time_manager/entity"
	"backend_time_manager/utils"
	"context"
	"github.com/google/uuid"
)

func FindProjectByUuid(id uuid.UUID) (entity.Project, error) {
	var project entity.Project
	err := Db.Get(&project, "SELECT * FROM TBL_PROJECTS WHERE id_project_uuid = $1", id)
	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func ListProjects(idOwner int64) ([]entity.Project, error) {
	var projects []entity.Project
	err := Db.Get(&projects, "SELECT * FROM TBL_PROJECTS WHERE id_owner = $1", idOwner)
	if err != nil {
		return projects, err
	}
	return projects, nil
}

func SaveProject(project entity.Project) (entity.Project, error) {
	if project.Id <= 0 {
		return insertProject(project)
	}
	return updateProject(project)
}

func insertProject(project entity.Project) (entity.Project, error) {
	var query = "INSERT INTO tbl_projects (dt_updated_at, nm_project, vl_color, id_owner) VALUES (current_timestamp, :nm_project, :vl_color, :id_owner)"
	result, err := Db.NamedExecContext(context.Background(), query, project)
	if err != nil {
		return entity.Project{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return entity.Project{}, err
	}

	err = Db.GetContext(context.Background(), &project, "SELECT * FROM TBL_PROJECTS WHERE id_project = $1", id)
	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func updateProject(project entity.Project) (entity.Project, error) {
	var query = "UPDATE tbl_projects set dt_updated_at = current_timestamp, nm_project = :nm_project, vl_color = :vl_color, id_owner = :id_owner WHERE id_project = :id_project"
	_, err := Db.NamedExecContext(context.Background(), query, project)
	if err != nil {
		return entity.Project{}, err
	}

	err = Db.GetContext(context.Background(), &project, "SELECT * FROM TBL_PROJECTS WHERE id_project = $1", project.Id)
	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func IncludeTagOnProject(project entity.Project, idsTag []int64) (entity.Project, []entity.Tag, error) {
	var idsTagsOnProject []int64
	err := Db.Get(&idsTagsOnProject, "SELECT id_tag FROM crz_projects_x_tags WHERE id_project = $1", project.Id)
	if err != nil {
		return project, nil, err
	}

	idsTagAdd := utils.FilterList(idsTag, func(id int64) bool {
		contains := false
		for _, tagId := range idsTagsOnProject {
			if id == tagId {
				contains = true
				break
			}
		}
		return !contains
	})

	var query = "INSERT INTO crz_projects_x_tags (id_project, id_tag) VALUES (:id_project,:id_tag)"
	for _, tagId := range idsTagAdd {
		_, err = Db.ExecContext(context.Background(), query, project.Id, tagId)
		if err != nil {
			return project, nil, err
		}
	}

	_, err = Db.ExecContext(context.Background(), "DELETE FROM crz_projects_x_tags WHERE !(id_tag in ($1))", idsTag)

	var tagsOnProject []entity.Tag
	err = Db.Get(&tagsOnProject, "SELECT * FROM crz_projects_x_tags WHERE id_project = $1", project.Id)

	if err != nil {
		return project, nil, err
	}

	return project, tagsOnProject, nil
}
