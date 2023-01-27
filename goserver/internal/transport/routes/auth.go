package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setAuthRoute(api *gin.RouterGroup) {
	auth_group := api.Group("/auth")

	auth_group.POST("/login", rest.MockResilt)    //AuthController_login
	auth_group.POST("/logout", rest.MockResilt)   //AuthController_logout
	auth_group.POST("/register", rest.MockResilt) //RegisterController_registerUser
	auth_group.GET("/captcha", rest.MockResilt)   //CaptchaController_getCaptcha
}
