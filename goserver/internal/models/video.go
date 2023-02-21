package models

type baseVideo struct {
	VideoId          string            `json:"videoId"`
	Title            string            `json:"title"`
	VideoThumbnails  []VideoThumbnail  `json:"videoThumbnails"`
	AuthorThumbnails []AuthorThumbnail `json:"authorThumbnails"`
	Author           string            `json:"author"`
	AuthorId         string            `json:"authorId"`
	ViewCount        int               `json:"viewCount"`
}

type commonVideo struct {
	Published      int    `json:"published"`
	PublishedText  string `json:"publishedText"`
	AuthorVerified bool   `json:"authorVerified"`
	Description    string `json:"description"`
	LikeCount      int    `json:"likeCount"`
	LengthSeconds  int    `json:"lengthSeconds"`
}

type RecommendedVideo struct {
	baseVideo
	AuthorUrl     string `json:"authorUrl"`
	LengthSeconds int    `json:"lengthSeconds"`
}

type VideoBasicInfo struct {
	baseVideo
	commonVideo
	Live               bool   `json:"live"`
	AuthorThumbnailUrl string `json:"authorThumbnailUrl"`
	DislikeCount       int    `json:"dislikeCount"`
	LengthString       string `json:"lengthString"`
}

type LegacyFormat struct {
	MimeType         string `json:"mimeType"`
	QualityLabel     string `json:"qualityLabel"`
	Bitrate          int    `json:"bitrate"`
	AudioBitrate     int    `json:"audioBitrate"`
	Itag             int    `json:"itag"`
	Url              string `json:"url"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps"`
	ProjectionType   string `json:"projectionType"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
	HasVideo         bool   `json:"hasVideo"`
	HasAudio         bool   `json:"hasAudio"`
	Container        string `json:"container"`
	Codecs           string `json:"codecs"`
	VideoCodec       string `json:"videoCodec"`
	AudioCodec       string `json:"audioCodec"`
	IsLive           bool   `json:"isLive"`
	IsHLS            bool   `json:"isHLS"`
	IsDashMPD        bool   `json:"isDashMPD"`
}

type Video struct {
	baseVideo
	commonVideo
	Type              string             `json:"type"`
	Storyboards       string             `json:"storyboards"`
	DescriptionHtml   string             `json:"descriptionHtml"`
	Keywords          []string           `json:"keywords"`
	Paid              bool               `json:"paid"`
	Premium           bool               `json:"premium"`
	IsFamilyFriendly  bool               `json:"isFamilyFriendly"`
	AllowedRegions    []string           `json:"allowedRegions"`
	Genre             string             `json:"genre"`
	GenreUrl          string             `json:"genreUrl"`
	AuthorId          string             `json:"authorId"`
	AuthorUrl         string             `json:"authorUrl"`
	SubCount          int                `json:"subCount"`
	AllowRatings      bool               `json:"allowRatings"`
	Rating            int                `json:"rating"`
	IsListed          bool               `json:"isListed"`
	LiveNow           bool               `json:"liveNow"`
	IsUpcoming        bool               `json:"isUpcoming"`
	AdaptiveFormats   []string           `json:"adaptiveFormats"`
	LegacyFormats     []LegacyFormat     `json:"legacyFormats"`
	Chapters          []Chapter          `json:"chapters"`
	Captions          string             `json:"live"`
	RecommendedVideos []RecommendedVideo `json:"recommendedVideos"`
	DashManifest      string             `json:"dashManifest"`
}

type VideoThumbnail struct {
	//'144p', '240p', '360p', '720p', '1080p', '1440p', '2160p'];
	Quality string `json:"quality"`
	Url     string `json:"url"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type VideoVisit struct {
	videoId         string
	progressSeconds int
	lengthSeconds   int
	/** Format date-time */
	lastVisit string
}

type VideoVisitDetails struct {
	videoDetails VideoBasicInfo
	VideoVisit
}

type UserprofileDetails struct {
	username     string
	profileImage string
	videoHistory []VideoVisitDetails
	/** Format date-time */
	registeredAt            string
	totalVideosCount        int
	totalTimeString         string
	subscribedChannelsCount int
}

type HistoryResponse struct {
	videos     []VideoVisitDetails
	videoCount int
}

type Popular struct {
	Videos []VideoBasicInfo `json:"videos"`
	/** Format date-time */
	UpdatedAt string `json:"updatedAt"`
}

type SubscriptionFeedResponse struct {
	VideoCount int              `json:"videoCount"`
	Videos     []VideoBasicInfo `json:"videos"`
	/** Format date-time */
	LastRefresh string `json:"lastRefresh"`
}
