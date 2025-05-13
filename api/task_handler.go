package api

import "github.com/gin-gonic/gin"

func ConfigureTaskApiRoutes(privateRouter *gin.RouterGroup) {
	group := privateRouter.Group("/tasks")
	group.POST("/", handleCreateTask)
	group.GET("/", handleListTasks)
	group.GET("/:uuid", handleReadTask)
	group.PUT("/:uuid", handleUpdateTask)
	group.DELETE("/:uuid", handleDeleteTask)
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
