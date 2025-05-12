package api

import (
	"backend_time_manager/database"
	"backend_time_manager/entity"
	"backend_time_manager/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ConfigurePublicAuthorizationApiRoutes(router *gin.RouterGroup) {
	router.POST("/auth/sign-in", handleSignIn)
	router.POST("/auth/sign-in/validate-token", handleConfirmSignIn)
}

func ConfigurePrivateAuthorizationApiRoutes(router *gin.RouterGroup) {
	router.POST("/auth/refresh", handleRefreshToken)
	router.DELETE("/auth/logout", handleSignOut)
}

func ValidateAndLoadToken(context *gin.Context) {
	token := context.GetHeader("Authorization")

	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token = strings.Replace(token, "Bearer ", "", 1)
	token = strings.Replace(token, "bearer ", "", 1)

	claim, err := utils.ParseToken(token)
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := database.FindUserById(claim.UserID)

	if err != nil {
		_ = context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	session, err := database.FindSessionByUuid(claim.SessionId)

	if err != nil {
		_ = context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	context.Set("logged_user", user)
	context.Set("logged_session", session)
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
	session := context.MustGet("logged_session").(*entity.Session)
	user := context.MustGet("logged_user").(*entity.User)

	var refreshToken string
	if err := context.BindJSON(refreshToken); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No refresh token provided"})
		return
	}

	refreshClaims, err := utils.ParseToken(refreshToken)

	if err != nil {
		_ = context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if refreshClaims.UserID != user.Id || refreshClaims.SessionId != session.Id {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid refresh token"})
		return
	}

	newAuthToken, err := utils.GenerateAccessToken(user.Id, session.Id)

	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": newAuthToken})
}

func handleSignOut(context *gin.Context) {
	sessionCtx, sessionExists := context.Get("logged_session")
	if !sessionExists {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	session := sessionCtx.(*entity.Session)

	err := database.DeleteSession(session.Id)

	if err != nil {
		_ = context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.Status(http.StatusOK)
}
