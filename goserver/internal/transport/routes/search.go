package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setSearchRoute(api *gin.RouterGroup) {
	search_group := api.Group("/search")

	{
		search_group.GET("/", rest.MockResilt)              //SearchController_search
		search_group.GET("/filters", rest.MockResilt)       //SearchController_getFilters
		search_group.GET("/continuation", rest.MockResilt)  //SearchController_searchContinuation
		search_group.GET("/dislikes/{id}", rest.MockResilt) //VideosController_getDislikes
	}
}
