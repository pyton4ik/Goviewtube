package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/models"
	"net/http"
)

func MockResilt(c *gin.Context) {
	//id := c.Param("id")

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album not found"})
}

type VideoThumbnailDto struct {
	//'144p', '240p', '360p', '720p', '1080p', '1440p', '2160p'];
	Quality string `json:"quality"`
	Url     string `json:"videoId"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

// albums slice to seed record album data.
var fake_film_tumbs = []models.VideoThumbnail{{
	Quality: "720p",
	Url:     "https://ucarecdn.com/a9ed82c6-14f1-462e-849e-3b007c2aae87",
	Width:   100,
	Height:  100,
},
}

func getFakeVideo(id string) models.VideoBasicInfo {
	var return_value models.VideoBasicInfo
	return_value.VideoId = id
	return_value.Title = "Title"
	return_value.PublishedText = "PublishedText"
	return_value.Author = "Autor"
	return_value.AuthorId = "AutorID"
	return_value.AuthorVerified = true
	return_value.AuthorThumbnailUrl = "https://ucarecdn.com/a9ed82c6-14f1-462e-849e-3b007c2aae87"
	return_value.VideoThumbnails = fake_film_tumbs
	return_value.Description = "Description"

	return_value.ViewCount = 1100
	return_value.LikeCount = 500
	return_value.DislikeCount = 13

	return_value.LengthSeconds = 130
	return return_value
}

// postAlbums adds an album from JSON received in the request body.
func GetPopular(c *gin.Context) {
	var res models.Popular
	res.Videos = append(res.Videos, getFakeVideo("1100"))
	res.Videos = append(res.Videos, getFakeVideo("2100"))
	c.IndentedJSON(http.StatusCreated, res)
}
