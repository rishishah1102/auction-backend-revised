package auth

import (
	"auction-backend/middlewares"
	"auction-backend/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// this route is for logging in the user  and getting a token to make requests
func LoginOtpController(c *gin.Context) {
	// fetching data from body
	var request schemas.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// generating jwt token
	token, err := middlewares.GenerateToken(request.Email)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error(),
			"message": "Unable to generate token for authentication",
		})
		return
	}

	// sending response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"token": token,
	})
}
