package api

import "github.com/gin-gonic/gin"

func ConfigureAuthorizationApiRoutes(router *gin.Engine) {
	router.POST("/sign-in", HandleSignIn)
	router.POST("/sign-out", HandleSignOut)
}

func HandleSignIn(c *gin.Context) {

}

func HandleSignOut(c *gin.Context) {

}
