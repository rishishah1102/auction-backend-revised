package auth

import (
	"auction-backend/middlewares"
	"auction-backend/schemas"
	"auction-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// this route is for logging in the user  and getting a token to make requests
func LoginOtpController(c *gin.Context) {
	// Reading logger
	logger, err := utils.ConfigLogger()
	if err != nil {
		zap.Must(zap.NewProduction()).Error(err.Error())
		return
	}

	// fetching data from body
	var request schemas.User
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// generating jwt token
	token, err := middlewares.GenerateToken(request.Email)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":   err.Error(),
			"message": "Unable to generate token for authentication",
		})
		return

	}

	// sending response
	logger.Info("Login successfully with " + request.Email)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"token":   token,
	})
}
