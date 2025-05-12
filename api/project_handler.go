package api

import "github.com/gin-gonic/gin"

func ConfigureProjectApiRoutes(privateRouter *gin.RouterGroup) {
	privateRouter.POST("/project", handleCreateProject)
	privateRouter.GET("/project/:uuid", handleReadProject)
	privateRouter.GET("/projects", handleListProjects)
	privateRouter.PUT("/project/:uuid", handleUpdateProject)
	privateRouter.DELETE("/project/:uuid", handleDeleteProject)
}

func handleCreateProject(context *gin.Context) {
}
func handleReadProject(context *gin.Context) {
}
func handleListProjects(context *gin.Context) {
}
func handleUpdateProject(context *gin.Context) {
}
func handleDeleteProject(context *gin.Context) {
}
