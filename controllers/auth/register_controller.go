package auth

import (
	"auction-backend/database"
	"auction-backend/schemas"
	"auction-backend/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// This route is for getting username and email from frontend and sending otp via email
func RegisterController(c *gin.Context) {
	// Reading logger
	logger, err := utils.ConfigLogger()
	if err != nil {
		zap.Must(zap.NewProduction()).Error(err.Error())
		return
	}

	// fetching data from body
	var request, response schemas.User
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check if user exists in db
	err = database.Users.FindOne(context.Background(), bson.M{"email": request.Email}).Decode(&response)
	if err == mongo.ErrNoDocuments {
		// generating opt
		otp := utils.GenerateRandomNumber()

		// sending email
		if err := utils.SendEmail(request.Email, "Registration Otp", otp); err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		// sending response
		logger.Info("OTP sent to the registered Email Id: " + request.Email)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Otp is send to " + request.Email,
			"otp":     otp,
		})
	} else if err == nil {
		logger.Warn("User already exist")
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
	} else {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

}
