package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setProfileRoute(api *gin.RouterGroup) {
	profile_group := api.Group("/profile")

	profile_group.GET("/details", rest.MockResilt)          //UserController_getProfileDetails
	profile_group.GET("/image/{username}", rest.MockResilt) //UserController_getProfileImage
	profile_group.POST("/image", rest.MockResilt)           //UserController_uploadProfileImage
	profile_group.DELETE("/image", rest.MockResilt)         //UserController_deleteProfileImage
	profile_group.POST("/password", rest.MockResilt)        //UserController_changePassword
}
