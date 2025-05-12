package api

import "github.com/gin-gonic/gin"

func ConfigureTagApiRoutes(privateRouter *gin.RouterGroup) {
	privateRouter.POST("/tag", handleCreateTag)
	privateRouter.GET("/tag/:id", handleReadTag)
	privateRouter.GET("/tags", handleListTags)
	privateRouter.PUT("/tag/:id", handleUpdateTag)
	privateRouter.DELETE("/tag/:id", handleDeleteTag)
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
