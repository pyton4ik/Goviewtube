package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setPlaylistRoute(api *gin.RouterGroup) {
	playlist_group := api.Group("/playlists")

	playlist_group.GET("/{playlistId}", rest.MockResilt) //PlaylistsController_getPlaylist
	playlist_group.GET("/continuation", rest.MockResilt) //PlaylistsController_getPlaylistContinuation

}
