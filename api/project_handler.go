package api

import "github.com/gin-gonic/gin"

func ConfigureProjectApiRoutes(router *gin.Engine) {
	router.POST("/project", handleCreateProject)
	router.GET("/project/:uuid", handleReadProject)
	router.GET("/projects", handleListProjects)
	router.PUT("/project/:uuid", handleUpdateProject)
	router.DELETE("/project/:uuid", handleDeleteProject)
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
