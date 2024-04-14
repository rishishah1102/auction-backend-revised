package models

import (
	"auction-backend/database"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func UsersCollection(logger *zap.Logger) (*mongo.Collection, error) {
	if database.DB == nil {
		logger.Warn("database is nil")
		return nil, errors.New("database client is nil")
	}

	usersCollection := database.DB.Collection("users")
	if usersCollection == nil {
		logger.Error("unable to make users collection")
		return nil, errors.New("unable to make users collection")
	}
	return usersCollection, nil
}
