package api

import (
	"backend_time_manager/database"
	"backend_time_manager/dto"
	"backend_time_manager/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ConfigureProjectApiRoutes(privateRouter *gin.RouterGroup) {
	privateRouter.POST("/project", handleCreateProject)
	privateRouter.GET("/projects", handleListProjects)
	routerWithProject := privateRouter.Group("/project/:projectId", loadProjectContext)
	routerWithProject.GET("/", handleReadProject)
	routerWithProject.PUT("/", handleUpdateProject)
	routerWithProject.DELETE("/", handleDeleteProject)
}

func handleCreateProject(context *gin.Context) {
	loggedUser := getLoggedUser(context)

	var projectForm dto.SaveProjectDto
	err := context.BindJSON(&projectForm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := entity.Project{
		IdOwner: loggedUser.Id,
		Name:    projectForm.Name,
		Color:   projectForm.Color,
	}

	project, err = database.SaveProject(project)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var idsTags []int64
	for _, tagDto := range projectForm.Tags {
		idsTags = append(idsTags, tagDto.Id)
	}

	project, tags, err := database.IncludeTagOnProject(project, idsTags)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	projectResponse := dto.ProjectDto{
		Id:        project.Uuid,
		Name:      project.Name,
		Color:     project.Color,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		Owner: dto.GenericEntityDto{
			Id:   loggedUser.Uuid,
			Name: loggedUser.Name,
		},
	}

	for _, tag := range tags {
		projectResponse.Tags = append(projectResponse.Tags, dto.TagDto{
			Id:    tag.Id,
			Name:  tag.Name,
			Color: tag.Color,
		})
	}

	context.JSON(http.StatusCreated, projectResponse)
}
func handleReadProject(context *gin.Context) {
	if context.IsAborted() {
		return
	}
	project := getProjectFromPath(context)
	loggedUser := getLoggedUser(context)

	if project.IdOwner != loggedUser.Id {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}

	projectDto := dto.ProjectDto{}.From(project)
	user, err := database.FindUserById(project.IdOwner)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
		return
	}
	projectDto.Owner = dto.GenericEntityDto{
		Id:   user.Uuid,
		Name: user.Name,
	}
	projectDto.Tags = []dto.TagDto{}

	// TODO load tags

	context.JSON(http.StatusOK, projectDto)
}
func handleListProjects(context *gin.Context) {
	loggedUser := getLoggedUser(context)

	projects, err := database.ListProjects(loggedUser.Id)
	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	projectDtos := make([]dto.ProjectDto, len(projects))
	for i, project := range projects {
		projectDtos[i] = dto.ProjectDto{}.From(project)
		projectDtos[i].Owner = dto.GenericEntityDto{
			Id:   loggedUser.Uuid,
			Name: loggedUser.Name,
		}
		projectDtos[i].Tags = []dto.TagDto{}

		// TODO load tags
	}

	context.JSON(http.StatusOK, projectDtos)
}
func handleUpdateProject(context *gin.Context) {
	project := getProjectFromPath(context)
	loggedUser := getLoggedUser(context)

	if project.IdOwner != loggedUser.Id {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}
}
func handleDeleteProject(context *gin.Context) {
	project := getProjectFromPath(context)
	loggedUser := getLoggedUser(context)

	if project.IdOwner != loggedUser.Id {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
		return
	}
}

func loadProjectContext(context *gin.Context) {
	id := context.Param("projectId")

	if err := uuid.Validate(id); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid project"})
		return
	}

	project, err := database.FindProjectByUuid(uuid.MustParse(id))
	if err != nil {
		context.String(http.StatusNotFound, "Project not found")
		return
	}

	context.Set("project", project)
}

func getProjectFromPath(context *gin.Context) entity.Project {
	return context.MustGet("project").(entity.Project)
}
