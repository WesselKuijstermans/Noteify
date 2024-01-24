package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("AuthMiddleWare")

		// Skip authentication for the following endpoints
		if strings.HasPrefix(c.FullPath(), "/auth") {
			c.Next()
			return
		}
		authHeader := c.GetHeader("Authorization")

		// Check if present
		if authHeader == "" {
			c.JSON(401, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Remove "Bearer " from the header
		token := authHeader[7:]

		valid, err := validateJWT(token)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "Unexpected signing method",
			})
			c.Abort()
			return
		}

		if !valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
		}

		// Logic
		c.Next()
	}
}

// check if the is jwt is from this api
func validateJWT(tokenString string) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env file")
	}
	secret := os.Getenv("JWT_SECRET")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// Token is malformed
				return false, nil
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return false, nil
			} else {
				// Token is invalid
				return false, nil
			}
		}
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// token is valid
		return true, nil
	}
	//	token is invalid
	return false, nil
}
