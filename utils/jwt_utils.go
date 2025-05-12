package utils

import (
	"backend_time_manager/entity"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

var accessSecret []byte
var refreshSecret []byte

func InitJwt() {
	accessSecret = []byte(os.Getenv("ACCESS_SECRET"))
	refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))
}

func GenerateAccessToken(userId int64, session entity.Session) (string, error) {
	expireAt := time.Now().Add(3 * time.Hour)
	if expireAt.After(session.ExpireAt) {
		expireAt = session.ExpireAt
	}

	accessClaims := KhiomisClaims{
		UserID:    userId,
		SessionId: session.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt), // access token valid for 15 min
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   session.Id.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	tokenString, err := token.SignedString(accessSecret)

	return tokenString, err
}

func GenerateRefreshToken(userId int64, session entity.Session) (string, error) {
	refreshClaims := KhiomisClaims{
		UserID:    userId,
		SessionId: session.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(session.ExpireAt), // access token valid for 15 min
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   session.Id.String(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(refreshSecret)

	return signedRefreshToken, err
}

func ParseToken(tokenString string) (*KhiomisClaims, error) {
	claims := &KhiomisClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure token method is what you expect (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})
	if err != nil {
		return nil, err
	}

	// Validate token and claims
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Optional: manually check expiration (usually handled automatically)
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return claims, fmt.Errorf("token expired")
	}

	return claims, nil
}

type KhiomisClaims struct {
	UserID    int64     `json:"user_id"`
	SessionId uuid.UUID `json:"session_id"`
	jwt.RegisteredClaims
}
