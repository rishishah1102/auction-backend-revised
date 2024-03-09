package add_all_player

import (
	"auction-backend/database"
	"auction-backend/schemas"
	"auction-backend/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// this route is for adding all the players into database which I get via converting a excel file.
func AddPlayerController(c *gin.Context) {
	// Reading logger
	logger, err := utils.ConfigLogger()
	if err != nil {
		zap.Must(zap.NewProduction()).Error(err.Error())
		return
	}

	// slice of players
	players := make([]schemas.Player, 250)

	// binding data from request body
	if err := c.ShouldBind(&players); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		for _, player := range players {
			// making a match model for every player
			var match schemas.Match
			matchData, err := database.Matches.InsertOne(context.Background(), match)
			if err != nil {
				logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				break
			}

			// adding _id of match model to each player
			insertedId, ok := matchData.InsertedID.(primitive.ObjectID)
			if !ok {
				logger.Error("Failed to convert _id to ObjectID")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Cannot convert _id to ObjectID",
				})
				break
			}
			player.Match = insertedId

			// inserting player into database
			_, err = database.Players.InsertOne(context.Background(), player)
			if err != nil {
				logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				break
			}
		}

		// sending response
		logger.Info("All players data inserted successfully")
		c.JSON(http.StatusOK, gin.H{
			"succes":  "Ok",
			"message": "Data inserted successfully",
		})
	}
}
