package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

// See /opt/Goviewtube/shared/api.schema.ts for source
func SetRoutes(api *gin.RouterGroup) {
	{
		api.GET("/videoplayback", rest.MockResilt)   //VideoplaybackController_getVideoplayback
		api.GET("/autocomplete", rest.MockResilt)    //AutocompleteController_getQuery
		api.GET("/dislikes/{id}", rest.MockResilt)   //VideosController_getDislikes
		api.GET("homepage/popular", rest.GetPopular) //HomepageController_getPopular
	}

	setAuthRoute(api)
	setHistoryRoute(api)
	setSubscriptionRoute(api)
	setProfileRoute(api)
	setUserRoute(api)
	setProxyRoute(api)
	setPlaylistRoute(api)
	setVideoRoute(api)
	setCommentRoute(api)
	setChannelRoute(api)
	setSearchRoute(api)
}
