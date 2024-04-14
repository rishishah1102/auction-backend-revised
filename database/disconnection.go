package database

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

func DisconnectDB(logger *zap.Logger) {
	// disconnecting the mongo
	if Client == nil {
		logger.Error("mongo client is nil", zap.Error(errors.New("mongo client is nil")))
		return
	}
	err := Client.Disconnect(context.Background())
	if err != nil {
		logger.Error("unable to disconnect with mongodb database", zap.Error(err))
		return
	}
	logger.Info("disconnected from database")
}
