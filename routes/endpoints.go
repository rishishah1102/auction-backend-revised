package routes

import (
	"auction-backend/controllers/add_all_player"
	"auction-backend/controllers/auth"
	"auction-backend/controllers/home"
	"auction-backend/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EndPoints(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Auction backend")
	})

	// making a group for authenticated endpoints
	authUsers := router.Group("/auth")
	authUsers.Use(middlewares.VerifyToken)

	// SIGNUP || METHOD POST
	router.POST("/signup", auth.RegisterController)

	// LOGIN || METHOD POST
	router.POST("/signin", auth.LoginController)

	// // OTP PAGE TO SAVE USER || METHOD POST
	router.POST("/otp", auth.RegisterOtpController)

	// // OTP PAGE TO LOGIN USER || METHOD POST
	router.POST("/loginotp", auth.LoginOtpController)

	// Player ADDING || METHOD POST
	authUsers.POST("/addplayer", add_all_player.AddPlayerController)

    // HOME || METHOD GET
	authUsers.GET("/home", home.HomeController)
}
