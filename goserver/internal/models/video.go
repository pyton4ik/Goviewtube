package models

type VideoThumbnailDto struct {
	quality string
	url     string
	width   int
	height  int
}

type RecommendedVideoDto struct {
	videoId          string
	title            string
	videoThumbnails  []VideoThumbnailDto
	author           string
	authorUrl        string
	authorId         string
	authorThumbnails []AuthorThumbnailDto
	lengthSeconds    int
	viewCountText    string
	viewCount        int
}

type VideoDto struct {
	Type              string
	title             string
	videoId           string
	videoThumbnails   []VideoThumbnailDto
	storyboards       string
	description       string
	descriptionHtml   string
	published         int
	publishedText     string
	keywords          []string
	viewCount         int
	likeCount         int
	paid              bool
	premium           bool
	isFamilyFriendly  bool
	allowedRegions    []string
	genre             string
	genreUrl          string
	author            string
	authorId          string
	authorUrl         string
	authorThumbnails  []AuthorThumbnailDto
	authorVerified    bool
	subCount          int
	lengthSeconds     int
	allowRatings      bool
	rating            int
	isListed          bool
	liveNow           bool
	isUpcoming        bool
	adaptiveFormats   []string
	legacyFormats     []string
	chapters          []ChapterDto
	captions          string
	recommendedVideos []RecommendedVideoDto
	dashManifest      string
}

type VideoVisitDto struct {
	videoId         string
	progressSeconds int
	lengthSeconds   int
	/** Format date-time */
	lastVisit string
}

type VideoVisitDetailsDto struct {
	videoDetails    VideoBasicInfoDto
	videoId         string
	progressSeconds int
	lengthSeconds   int
	/** Format date-time */
	lastVisit string
}

type UserprofileDetailsDto struct {
	username     string
	profileImage string
	videoHistory []VideoVisitDetailsDto
	/** Format date-time */
	registeredAt            string
	totalVideosCount        int
	totalTimeString         string
	subscribedChannelsCount int
}

type HistoryResponseDto struct {
	videos     []VideoVisitDetailsDto
	videoCount int
}

type VideoBasicInfoDto struct {
	videoId            string
	title              string
	published          int
	publishedText      string
	author             string
	authorId           string
	authorVerified     bool
	authorThumbnails   []AuthorThumbnailDto
	authorThumbnailUrl string
	videoThumbnails    []VideoThumbnailDto
	description        string
	viewCount          int
	likeCount          int
	dislikeCount       int
	lengthSeconds      int
	lengthString       string
	live               bool
}

type PopularDto struct {
	videos []VideoBasicInfoDto
	/** Format date-time */
	updatedAt string
}

type SubscriptionFeedResponseDto struct {
	videoCount int
	videos     []VideoBasicInfoDto
	/** Format date-time */
	lastRefresh string
}
