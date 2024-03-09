package add_all_player

import (
	"auction-backend/database"
	"auction-backend/schemas"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this route is for adding all the players into database which I get via converting a excel file.
func AddPlayerController(c *gin.Context) {
	// slice of players
	players := make([]schemas.Player, 250)

	// binding data from request body
	if err := c.ShouldBind(&players); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		for _, player := range players {
			// making a match model for every player
			var match schemas.Match
			matchData, err := database.Matches.InsertOne(context.Background(), match)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				break
			}

			// adding _id of match model to each player
			insertedId, ok := matchData.InsertedID.(primitive.ObjectID)
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Cannot convert _id to ObjectID",
				})
				break
			}
			player.Match = insertedId

			// inserting player into database
			_, err = database.Players.InsertOne(context.Background(), player)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				break
			}
		}

		// sending response
		c.JSON(http.StatusOK, gin.H{
			"succes":  "Ok",
			"message": "Data inserted successfully",
		})
	}
}
