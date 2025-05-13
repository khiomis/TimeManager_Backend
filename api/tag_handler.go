package api

import "github.com/gin-gonic/gin"

func ConfigureTagApiRoutes(privateRouter *gin.RouterGroup) {
	group := privateRouter.Group("/tags")
	group.POST("/", handleCreateTag)
	group.GET("/", handleListTags)
	group.GET("/:id", handleReadTag)
	group.PUT("/:id", handleUpdateTag)
	group.DELETE("/:id", handleDeleteTag)
}

func handleCreateTag(context *gin.Context) {
}
func handleReadTag(context *gin.Context) {
}
func handleListTags(context *gin.Context) {
}
func handleUpdateTag(context *gin.Context) {
}
func handleDeleteTag(context *gin.Context) {
}
