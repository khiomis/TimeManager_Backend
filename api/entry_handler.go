package api

import "github.com/gin-gonic/gin"

func ConfigureEntryApiRoutes(privateRouter *gin.RouterGroup) {
	group := privateRouter.Group("/entry")
	group.POST("/", handleCreateEntry)
	group.GET("/:uuid", handleReadEntry)
	group.GET("/list", handleListEntries)
	group.PUT("/:uuid", handleUpdateEntry)
	group.DELETE("/:uuid", handleDeleteEntry)
}

func handleCreateEntry(context *gin.Context) {
}
func handleReadEntry(context *gin.Context) {
}
func handleListEntries(context *gin.Context) {
}
func handleUpdateEntry(context *gin.Context) {
}
func handleDeleteEntry(context *gin.Context) {
}
