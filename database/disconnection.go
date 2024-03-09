package database

import (
    "context"
    "go.uber.org/zap"
)

func DisconnectDB(logger *zap.Logger) {
    // disconnecting the mongo
    if Client == nil {
        logger.Error("Mongo Client is nil")
        return
    }
    err := Client.Disconnect(context.Background())
    if err != nil {
        logger.Error(err.Error())
        return
    }
    logger.Info("Disconnected from Database")
}