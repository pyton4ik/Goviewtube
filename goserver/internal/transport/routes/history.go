package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setHistoryRoute(api *gin.RouterGroup) {
	history_group := api.Group("/history")

	history_group.GET("/videos", rest.MockResilt)                           //HistoryController_getHistory
	history_group.DELETE("/{channelId}", rest.MockResilt)                   //HistoryController_deleteEntireHistory
	history_group.GET("/{id}", rest.MockResilt)                             //HistoryController_getVideoVisit
	history_group.POST("/{id}", rest.MockResilt)                            //HistoryController_setVideoVisit
	history_group.DELETE("/{videoId}", rest.MockResilt)                     //HistoryController_deleteHistoryEntry
	history_group.DELETE("/from/{startDate}/to/{endDate}", rest.MockResilt) //HistoryController_deleteHistoryRange
}
