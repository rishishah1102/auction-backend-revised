package home

import (
	"auction-backend/database"
	"auction-backend/schemas"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// this route is for fetching the image and username for the home page
func HomeController(c *gin.Context) {
	// fetching email from token
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email not found in context",
		})
		return
	}

	// finding  user by email
	var result schemas.User
	err := database.Users.FindOne(context.Background(), bson.M{"email": email}).Decode(&result)
	if err != nil {
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
