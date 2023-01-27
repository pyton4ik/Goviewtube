package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setSubscriptionRoute(api *gin.RouterGroup) {
	subscription_group := api.Group("/subscriptions")

	subscription_group.GET("/channels", rest.MockResilt)       //SubscriptionsController_getSubscribedChannels
	subscription_group.GET("/videos", rest.MockResilt)         //SubscriptionsController_getSubscriptionVideos
	subscription_group.GET("/{channelId}", rest.MockResilt)    //SubscriptionsController_getSubscription
	subscription_group.PUT("/{channelId}", rest.MockResilt)    //SubscriptionsController_createSubscription
	subscription_group.DELETE("/{channelId}", rest.MockResilt) //SubscriptionsController_deleteSubscription
	subscription_group.POST("/multiple", rest.MockResilt)      //SubscriptionsController_createMultipleSubscriptions
}
