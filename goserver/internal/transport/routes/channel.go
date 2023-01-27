package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setChannelRoute(api *gin.RouterGroup) {
	channels_group := api.Group("/channels")
	{
		channels_group.GET("{id}/thumbnail/tiny.jpg", rest.MockResilt)       //ChannelsController_getTinyThumbnailJpg
		channels_group.GET("{id}/thumbnail/tiny.webp", rest.MockResilt)      //ChannelsController_getTinyThumbnailWebp
		channels_group.GET("/{id}", rest.MockResilt)                         //ChannelsController_getChannelInfo
		channels_group.GET("/{id}/home", rest.MockResilt)                    //ChannelsController_getChannelHome
		channels_group.GET("/{id}/videos", rest.MockResilt)                  //ChannelsController_getChannelVideos
		channels_group.GET("/videos/continuation", rest.MockResilt)          //ChannelsController_getChannelVideosContinuation
		channels_group.GET("/{id}/playlists", rest.MockResilt)               //ChannelsController_getChannelPlaylists
		channels_group.GET("/playlists/continuation", rest.MockResilt)       //ChannelsController_getChannelPlaylistsContinuation
		channels_group.GET("{id}/search", rest.MockResilt)                   //ChannelsController_searchChannel
		channels_group.GET("/search/continuation", rest.MockResilt)          //ChannelsController_searchChannelContinuation
		channels_group.GET("/relatedchannels/continuation", rest.MockResilt) //ChannelsController_getRelatedChannelsContinuation
		channels_group.GET("/{id}/communityposts", rest.MockResilt)          //ChannelsController_getChannelCommunityPosts
		channels_group.GET("/communityposts/continuation", rest.MockResilt)  //ChannelsController_getChannelCommunityPostsContinuation
		channels_group.GET("/{id}/stats", rest.MockResilt)                   //ChannelsController_getChannelStats
	}
}
