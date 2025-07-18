package middleware

import (
	"strconv"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	userSecrets map[string]string
	mutex       sync.RWMutex
}

var tokenManager = &TokenManager{userSecrets: make(map[string]string)}

func generateUserSecret(userID string) string {
	return userID + "_" + time.Now().Format("20060102150405")
}

func GetUserSecret(userID string) string {
	tokenManager.mutex.RLock()
	secret, exists := tokenManager.userSecrets[userID]
	tokenManager.mutex.RUnlock()

	if exists {
		return secret
	}

	newSecret := generateUserSecret(userID)
	SetUserSecret(userID, newSecret)
	return newSecret
}

func SetUserSecret(userID string, secret string) {
	tokenManager.mutex.Lock()
	tokenManager.userSecrets[userID] = secret
	tokenManager.mutex.Unlock()
}

func InvalidateUserTokens(userID string) {
	tokenManager.mutex.Lock()
	tokenManager.userSecrets[userID] = generateUserSecret(userID)
	tokenManager.mutex.Unlock()
}

func GenerateToken(userID int64, username string) (string, error) {
	userIDStr := strconv.FormatInt(userID, 10)
	secret := GetUserSecret(userIDStr)

	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString string, userID int64) (*jwt.Token, error) {

	userIDStr := strconv.FormatInt(userID, 10)
	secret := GetUserSecret(userIDStr)

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
}
