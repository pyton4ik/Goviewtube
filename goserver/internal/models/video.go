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

type Video struct {
	baseVideo
	commonVideo
	Type              string             `json:"live"`
	storyboards       string             `json:"live"`
	descriptionHtml   string             `json:"live"`
	keywords          []string           `json:"live"`
	paid              bool               `json:"live"`
	premium           bool               `json:"live"`
	isFamilyFriendly  bool               `json:"live"`
	allowedRegions    []string           `json:"live"`
	genre             string             `json:"live"`
	genreUrl          string             `json:"live"`
	authorId          string             `json:"live"`
	authorUrl         string             `json:"live"`
	subCount          int                `json:"live"`
	allowRatings      bool               `json:"live"`
	rating            int                `json:"live"`
	isListed          bool               `json:"live"`
	liveNow           bool               `json:"live"`
	isUpcoming        bool               `json:"live"`
	adaptiveFormats   []string           `json:"live"`
	legacyFormats     []string           `json:"live"`
	chapters          []Chapter          `json:"live"`
	captions          string             `json:"live"`
	recommendedVideos []RecommendedVideo `json:"live"`
	dashManifest      string             `json:"live"`
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
