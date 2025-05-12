package api

import "github.com/gin-gonic/gin"

func ConfigureTaskApiRoutes(privateRouter *gin.RouterGroup) {
	privateRouter.POST("/task", handleCreateTask)
	privateRouter.GET("/task/:uuid", handleReadTask)
	privateRouter.GET("/tasks", handleListTasks)
	privateRouter.PUT("/task/:uuid", handleUpdateTask)
	privateRouter.DELETE("/task/:uuid", handleDeleteTask)
}

func handleCreateTask(context *gin.Context) {
}
func handleReadTask(context *gin.Context) {
}
func handleListTasks(context *gin.Context) {
}
func handleUpdateTask(context *gin.Context) {
}
func handleDeleteTask(context *gin.Context) {
}
