package home

import (
	"auction-backend/models"
	"auction-backend/schemas"
	"auction-backend/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

// this route is for fetching the image and username for the home page
func HomeController(c *gin.Context) {
	// Reading logger
	logger, err := utils.ConfigLogger()
	if err != nil {
		zap.Must(zap.NewProduction()).Error(err.Error())
		return
	}

	// fetching email from token
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email not found in context",
		})
		return
	}

	// fetching the user collection
	usersCollection, err := models.UsersCollection(logger)
	if err != nil {
		logger.Error("unable to get users collection", zap.Error(err))
		return
	}

	// finding  user by email
	var result schemas.User
	err = usersCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		logger.Error("unable to fetch the data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in fetching the data",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Successfully fetched the data",
			"foundUser": result,
		})
	}
}
