package api

import "github.com/gin-gonic/gin"

func ConfigureTagApiRoutes(router *gin.Engine) {
	router.POST("/tag", handleCreateTag)
	router.GET("/tag/:id", handleReadTag)
	router.GET("/tags", handleListTags)
	router.PUT("/tag/:id", handleUpdateTag)
	router.DELETE("/tag/:id", handleDeleteTag)
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
