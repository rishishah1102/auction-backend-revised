package add_all_player

import (
	"auction-backend/database"
	"auction-backend/schemas"
	"auction-backend/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
			matchData, err := database.Matches.InsertOne(context.Background(), bson.M{
				"match1":            match.Match1,
				"match2":            match.Match2,
				"match3":            match.Match3,
				"match4":            match.Match4,
				"match5":            match.Match5,
				"match6":            match.Match6,
				"match7":            match.Match7,
				"match8":            match.Match8,
				"match9":            match.Match9,
				"match10":           match.Match10,
				"prevX1":            match.PrevX1,
				"currentX1":         match.CurrentX1,
				"nextX1":            match.NextX1,
				"earnedPoints":      match.EarnedPoints,
				"benchedPoints":     match.BenchedPoints,
				"totalPoints":       match.TotalPoints,
				"prevTotalPoints":   match.PrevTotalPoints,
				"prevEarnedPoints":  match.PrevEarnedPoints,
				"prevBenchedPoints": match.PrevBenchedPoints,
			})
			if err != nil {
				logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			// adding _id of match model to each player
			insertedId, ok := matchData.InsertedID.(primitive.ObjectID)
			if !ok {
				logger.Error("Failed to convert _id to ObjectID")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Cannot convert _id to ObjectID",
				})
				return
			}
			player.Match = insertedId

			// inserting player into database
			_, err = database.Players.InsertOne(context.Background(), bson.M{
				"playerNumber":         player.PlayerNumber,
				"playerName":           player.PlayerName,
				"country":              player.Country,
				"playerType":           player.PlayerType,
				"iplTeam":              player.IplTeam,
				"prevTeam":             player.PrevTeam,
				"currentTeam":          player.CurrentTeam,
				"basePrice":            player.BasePrice,
				"prevFantasyPoints":    player.PrevFantasyPoints,
				"currentFantasyPoints": player.CurrentFantasyPoints,
				"sellingPrice":         player.SellingPrice,
				"match":                player.Match,
				"sold":                 player.Sold,
				"unsold":               player.Unsold,
			})
			if err != nil {
				logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
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
