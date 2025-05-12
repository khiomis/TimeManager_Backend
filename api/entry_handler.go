package api

import "github.com/gin-gonic/gin"

func ConfigureEntryApiRoutes(router *gin.Engine) {
	router.POST("/entry", handleCreateEntry)
	router.GET("/entry/:uuid", handleReadEntry)
	router.GET("/entries", handleListEntries)
	router.PUT("/entry/:uuid", handleUpdateEntry)
	router.DELETE("/entry/:uuid", handleDeleteEntry)
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
