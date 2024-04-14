package models

import (
	"auction-backend/database"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func PlayersCollection(logger *zap.Logger) (*mongo.Collection, error) {
	if database.DB == nil {
		logger.Warn("database is nil")
		return nil, errors.New("database client is nil")
	}

	playersCollection := database.DB.Collection("players")
	if playersCollection == nil {
		logger.Error("unable to make players collection")
		return nil, errors.New("unable to make players collection")
	}
	return playersCollection, nil
}
