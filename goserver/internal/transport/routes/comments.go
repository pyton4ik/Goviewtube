package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setCommentRoute(api *gin.RouterGroup) {
	comments_group := api.Group("/comments")

	comments_group.GET("/{videoId}", rest.MockResilt)         //CommentsController_getComments
	comments_group.GET("/{videoId}/replies", rest.MockResilt) //CommentsController_getCommentReplies

}
