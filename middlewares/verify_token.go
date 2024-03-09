package middlewares

import (
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(c *gin.Context) {
	// jwt Secret
	jwtKey := []byte(os.Getenv("TOKEN_SECRET"))

	// fetching token from header of request
	headerToken := c.Request.Header.Get("Authorization")
	if headerToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "A token is required for authentication",
		})
		return
	}

	// Parse the token
	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		// checking the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	// validating token
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.Abort()
		return
	}

	// fetching claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse token claims",
		})
		c.Abort()
		return
	}

	// extracting email from token claims
	email, ok := claims["email"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to extract email from token claims",
		})
		c.Abort()
		return
	}

	// Store the email in the context for further use
	c.Set("email", email)

	// if token is valid forwarding request
	c.Next()
}
