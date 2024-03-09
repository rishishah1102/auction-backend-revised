package database

import (
	"auction-backend/utils"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var Client *mongo.Client
var Users *mongo.Collection
var Players *mongo.Collection
var Matches *mongo.Collection

func ConnectDB(logger *zap.Logger, mongoConfig utils.MongoConfig) (error){
	// connection with mongo Atlas
	clientOptions := options.Client().ApplyURI(mongoConfig.MongoUri)
	cl, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	Client = cl

	// it checks whether the connection with MongoDb is established or not
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	// Creating database
	db := Client.Database(mongoConfig.Database)
	if db == nil {
		logger.Error("Unable to make database")
		return err
	}

	// Creating collections
	userCollection := db.Collection("users")
	if userCollection == nil {
		logger.Error("Unable to make users collection")
		return err
	}
	playersCollection := db.Collection("players")
	if playersCollection == nil {
		logger.Error("Unable to make players collection")
		return err
	}
	matchesCollection := db.Collection("matches")
	if matchesCollection == nil {
		logger.Error("Unable to make matches collection")
		return err
	}

	// exporting the collections
	Users = userCollection
	Players = playersCollection
	Matches = matchesCollection

	logger.Info("Connected to Database")
	return nil
}
