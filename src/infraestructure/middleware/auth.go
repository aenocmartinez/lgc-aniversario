package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || claims["user_id"] == nil {
				return nil, jwt.ErrInvalidKey
			}

			userID := claims["user_id"].(float64)
			userIDStr := strconv.FormatInt(int64(userID), 10)

			secret := GetUserSecret(userIDStr)

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok || time.Now().Unix() > int64(exp) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
			c.Abort()
			return
		}

		c.Set("user_id", int64(claims["user_id"].(float64)))

		c.Next()
	}
}
