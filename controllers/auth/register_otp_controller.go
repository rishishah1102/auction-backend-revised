package auth

import (
	"auction-backend/database"
	"auction-backend/middlewares"
	"auction-backend/schemas"
	"auction-backend/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func RegisterOtpController(c *gin.Context) {
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

	fmt.Println(request.Squad)

	// saving into the database
	_, err = database.Users.InsertOne(context.Background(), bson.M{"email": request.Email, "username": request.Username, "ImgUrl": request.ImgUrl, "teamname": request.Teamname, "squad": request.Squad, "isPlaying": request.IsPlaying, "isAuctioneer": request.IsAuctioneer})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusFound, gin.H{
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
	logger.Info("Registration successfully with " + request.Email)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration Successfully",
		"token":   token,
	})
}
