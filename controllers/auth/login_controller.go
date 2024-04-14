package auth

import (
	"auction-backend/models"
	"auction-backend/schemas"
	"auction-backend/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// This route logs in the user. It takes the email from user and sends otp
func LoginController(c *gin.Context) {
	// Reading logger
	logger, err := utils.ConfigLogger()
	if err != nil {
		zap.Must(zap.NewProduction()).Error(err.Error())
		return
	}

	// fetching data from body
	var request, response schemas.User
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("unable to bind the request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fetching the user collection
	usersCollection, err := models.UsersCollection(logger)
	if err != nil {
		logger.Error("unable to get users collection", zap.Error(err))
		return
	}

	// find user in database
	err = usersCollection.FindOne(context.Background(), bson.M{"email": request.Email}).Decode(&response)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Error("No user found with provided email", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{
				"message": "No user found with the provided email",
			})
		} else {
			logger.Error("error in fetching the data", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error in fetching the data",
				"error":   err.Error(),
			})
		}
		return
	}

	// generating opt
	otp := utils.GenerateRandomNumber()
	logger.Info("otp generated successfully")

	// sending email
	if err := utils.SendEmail(request.Email, "Login Otp", otp); err != nil {
		logger.Error("unable to send email", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// sending response
	logger.Info("otp sent to " + request.Email)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Otp is send to " + request.Email,
		"otp":     otp,
	})
}
