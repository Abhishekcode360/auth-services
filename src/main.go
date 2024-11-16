package main

import (
	jwtauthController "auth-services/controller/jwtauthController"
	refreshController "auth-services/controller/refreshController"
	revokeController "auth-services/controller/revokeController"
	signInController "auth-services/controller/signInController"
	signupController "auth-services/controller/signupController"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the server....")
	router := gin.New()

	router.POST("/signup", signupController.SignUp)
	router.POST("/signin", signInController.SignIn)
	router.GET("/tokenauth", jwtauthController.AuthenticateJWT)
	router.POST("/revoketoken", revokeController.Revoke)
	router.POST("/refresh", refreshController.Refresh)

	router.Run(":3001")
	fmt.Println("Exiting....")
}
