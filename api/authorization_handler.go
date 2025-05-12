package api

import "github.com/gin-gonic/gin"

func ConfigureAuthorizationApiRoutes(router *gin.Engine) {
	router.POST("/auth/sign-in", handleSignIn)
	router.POST("/auth/sign-in/validate-token", handleConfirmSignIn)
	router.POST("/auth/refresh", handleRefreshToken)
	router.DELETE("/auth/logout", handleSignOut)
}

func handleSignIn(context *gin.Context) {
	// Get the sign in data from the request
	// Validate
	// Generate a token and send to the user's email
	// Return success with a few data needed to be sent after to validate the token
}

func handleConfirmSignIn(context *gin.Context) {
	// Get token from body
	// Validate token, checking all data available
	// Create a session and a jwt token with the session id and the user id
	// Return success, with a few infos of the user, the jwt token and refresh token
}

func handleRefreshToken(context *gin.Context) {
	// Get the refresh token
	// Generate the new auth token
	// Return the generated auth token
}

func handleSignOut(context *gin.Context) {
	// Get if the auth is valid
	// Get the session id and user id from the jwt token
	// Delete the session from the db
	// Return success
}
