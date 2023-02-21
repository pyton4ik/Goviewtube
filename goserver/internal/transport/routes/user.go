package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setUserRoute(api *gin.RouterGroup) {
	user_group := api.Group("/user")
	{
		user_group.GET("/subscriptions/videos", rest.GetSubVideos) //SubscriptionsController_getSubscriptionVideos
		user_group.GET("/profile", rest.MockResilt)                //UserController_getProfile
		user_group.GET("/export", rest.MockResilt)                 //UserController_getExport

		user_group.GET("/settings", rest.MockResilt)                 //SettingsController_getSettings
		user_group.PUT("/settings", rest.MockResilt)                 //SettingsController_setSettings
		user_group.GET("/", rest.MockResilt)                         //UserController_getProfile
		user_group.POST("/notifications/subscribe", rest.MockResilt) //NotificationsController_subscribeToNotifications

	}

}
