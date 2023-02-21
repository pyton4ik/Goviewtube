package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/models"
	"net/http"
)

func GetChannelInfo(c *gin.Context) {
	var res models.ChannelInfo
	id := c.Query("id")
	fmt.Println("dummy GetChannelInfo with id", id)
	c.IndentedJSON(http.StatusCreated, res)
}

func GetChannelHome(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("dummy GetChannelHome with id", id)
	var res models.ChannelHome
	c.IndentedJSON(http.StatusCreated, res)

}

func GetDashManifest(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("dummy GetChannelHome with id", id)
	c.String(http.StatusOK, "Hello fake %s ", id)
}
