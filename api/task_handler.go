package api

import "github.com/gin-gonic/gin"

func ConfigureTaskApiRoutes(router *gin.Engine) {
	router.POST("/task", handleCreateTask)
	router.GET("/task/:uuid", handleReadTask)
	router.GET("/tasks", handleListTasks)
	router.PUT("/task/:uuid", handleUpdateTask)
	router.DELETE("/task/:uuid", handleDeleteTask)
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
