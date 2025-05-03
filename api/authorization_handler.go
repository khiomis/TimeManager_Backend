package api

import "github.com/gin-gonic/gin"

func ConfigureAuthorizationApiRoutes(router *gin.Engine) {
	router.POST("/auth", handleSignIn)
	router.PUT("/auth", handleRefreshToken)
	router.DELETE("/auth", handleSignOut)
}

func handleSignIn(c *gin.Context) {

}

func handleRefreshToken(c *gin.Context) {

}

func handleSignOut(c *gin.Context) {

}
