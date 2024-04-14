package handler

import (
	"auction-backend/database"
	"auction-backend/routes"
	"auction-backend/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	Router *gin.Engine
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func init() {
	gin.SetMode(gin.ReleaseMode)

	// Reading yaml file
	logger, err := utils.ConfigLogger()
	if err != nil {
		zap.Must(zap.NewProduction()).Error("failed to initialize custom logger", zap.Error(err))
		return
	}

	// mongo config
	var mongoConfig utils.MongoConfig
	mongoConfig.MongoUri = os.Getenv("MONGODB_URI")
	mongoConfig.Database = os.Getenv("DATABASE_NAME")

	// gin instance
	router := gin.Default()
	Router = router

	// cors
	router.Use(CORSMiddleware())

	// Connecting with database
	err = database.ConnectDB(logger, mongoConfig)
	if err != nil {
		logger.Error("unable to connect with mongodb database", zap.Error(err))
		return
	}
	defer database.DisconnectDB(logger)

	//user_routes
	routes.EndPoints(router)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// serverless
	Router.ServeHTTP(w, r)
}
