package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/rest"
)

func setProxyRoute(api *gin.RouterGroup) {
	proxy_group := api.Group("/proxy")

	proxy_group.GET("/text", rest.MockResilt)   //ProxyController_getText
	proxy_group.GET("/image", rest.MockResilt)  //ProxyController_getQuery
	proxy_group.GET("/stream", rest.MockResilt) //ProxyController_proxyStream

}
