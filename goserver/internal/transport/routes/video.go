package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setVideoRoute(api *gin.RouterGroup) {
	video_group := api.Group("/videos")
	{
		video_group.GET("/{id}", rest.MockResilt)               //VideosController_getVideos
		video_group.GET("/manifest/dash/{id}", rest.MockResilt) //VideosController_getDashManifest
		video_group.GET("/dislikes/{id}", rest.MockResilt)      //VideosController_getDislikes
	}

}
