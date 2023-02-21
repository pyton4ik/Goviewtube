package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/pyton4ik/Goviewtube/goserver/internal/models"
	"net/http"
)

func GetVideoByID(c *gin.Context) {
	var res models.Video
	var format models.LegacyFormat
	format.MimeType = "video/mp4; codecs=\"avc1.64001F, mp4a.40.2\""
	format.QualityLabel = "720p"
	format.Bitrate = 1189003
	format.AudioBitrate = 192
	format.Itag = 22
	format.Url = "https://rr2---sn-4g5lzner.googlevideo.com/videoplayback?expire=1676982967&ei=Vmb0Y6GKO86ox_APidOdoA0&ip=156.238.9.118&id=o-AKeVlAj0IcMB9M9RnlD46w-tMZ9lUACRLq_yJT36Qovi&itag=22&source=youtube&requiressl=yes&mh=0K&mm=31%2C29&mn=sn-4g5lzner%2Csn-4g5ednds&ms=au%2Crdu&mv=m&mvi=2&pl=23&initcwndbps=217500&vprv=1&mime=video%2Fmp4&ns=GR5yd38B9hzFaQjfkN2tSYkL&cnr=14&ratebypass=yes&dur=755.670&lmt=1676160033828458&mt=1676960939&fvip=3&fexp=24007246&c=WEB&txp=4532434&n=mAy-3eZdTSIwow&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cns%2Ccnr%2Cratebypass%2Cdur%2Clmt&sig=AOq0QJ8wRAIgda4RHmXlT7OxgnLnJ6vuFJay2SQYQV99MmJ2WUFrnqkCIHgH6c-q7byUfc0V9DdJOLbntGPTBlL_nCaxmHW2MwM_&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AG3C_xAwRAIgHbjA51fI_sPsOI6h-n09bjeAjWgCL-r529nFN6erz2gCIAhfwX0vRpOThSZKVJqq7V7ab2VTcRvJ655UIOD2SY_S"
	format.Width = 1280
	format.Height = 720
	format.LastModified = "1676160033828458"
	format.Quality = "hd720"
	format.Fps = 30
	format.ProjectionType = "RECTANGULAR"
	format.AudioQuality = "AUDIO_QUALITY_MEDIUM"
	format.ApproxDurationMs = "755670"
	format.AudioSampleRate = "44100"
	format.AudioChannels = 2
	format.HasVideo = true
	format.HasAudio = true
	format.Container = "mp4"
	format.Codecs = "avc1.64001F, mp4a.40.2"
	format.VideoCodec = "avc1.64001F"
	format.AudioCodec = "mp4a.40.2"
	format.IsLive = false
	format.IsHLS = false
	format.IsDashMPD = false

	//id := "XXDDDFFF"
	id := c.Query("id")

	res.VideoId = id
	res.Title = "Title" + id
	res.PublishedText = "PublishedText" + id
	res.Author = "Autor" + id
	res.AuthorId = "AutorID" + id
	res.AuthorVerified = true
	//res.AuthorThumbnailUrl = "https://ucarecdn.com/a9ed82c6-14f1-462e-849e-3b007c2aae87"
	res.VideoThumbnails = fake_film_tumbs
	res.Description = "Description"

	res.ViewCount = 1100
	res.LikeCount = 55090
	//res.DislikeCount = 13

	res.LengthSeconds = 130

	res.Type = "type"
	res.Storyboards = "storyboards"
	res.DescriptionHtml = "<title>HTML base Tag</title>"
	res.Keywords = []string{"keyword1", "keyword2"}
	res.Paid = true
	res.Premium = true
	res.IsFamilyFriendly = true
	res.AllowedRegions = []string{"ALLOWED_REGION1", "ALLOWED_REGION2"}
	res.Genre = "genre"
	res.AuthorUrl = "http://authorUrl"
	res.SubCount = 456455
	res.AllowRatings = true
	res.Rating = 5456
	res.IsListed = true
	res.LiveNow = false
	res.IsUpcoming = true
	res.AdaptiveFormats = []string{}
	res.Chapters = []models.Chapter{}
	res.Captions = "captions"
	res.RecommendedVideos = []models.RecommendedVideo{}
	res.DashManifest = "dashManifest"
	res.LegacyFormats = append(res.LegacyFormats, format)

	c.IndentedJSON(http.StatusCreated, res)
}

func GetDislikes(c *gin.Context) {
	id := c.Query("id")
	var res = models.Dislike{Id: id,
		Likes:       345,
		DateCreated: "",
		Dislikes:    444,
		Rating:      6766,
		ViewCount:   210349,
		Deleted:     false,
	}
	c.IndentedJSON(http.StatusCreated, res)
}
