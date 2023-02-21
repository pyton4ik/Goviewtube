package models

type ChannelImage struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type RelatedChannel struct {
	ChannelName     string         `json:"channelName"`
	ChannelId       string         `json:"channelId"`
	ChannelUrl      string         `json:"channelUrl"`
	Thumbnail       []ChannelImage `json:"thumbnail"`
	VideoCount      int            `json:"videoCount"`
	SubscriberText  string         `json:"subscriberText"`
	SubscriberCount int            `json:"subscriberCount"`
	Verified        bool           `json:"verified"`
	OfficialArtist  bool           `json:"officialArtist"`
	OfficialArist   bool           `json:"officialArist"`
}

type ChannelLink struct {
	Url   string `json:"url"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

type RelatedChannels struct {
	Items        []RelatedChannel `json:"items"`
	Continuation string           `json:"continuation"`
}

type ChannelLinks struct {
	primaryLinks   []ChannelLink `json:"published"`
	secondaryLinks []ChannelLink `json:"published"`
}

type ChannelInfo struct {
	author           string            `json:"published"`
	authorId         string            `json:"published"`
	authorUrl        string            `json:"published"`
	authorBanners    []ChannelImage    `json:"published"`
	authorThumbnails []ChannelImage    `json:"published"`
	subscriberText   string            `json:"published"`
	subscriberCount  int               `json:"published"`
	description      string            `json:"published"`
	isFamilyFriendly bool              `json:"published"`
	relatedChannels  []RelatedChannels `json:"published"`
	allowedRegions   []string          `json:"published"`
	isVerified       bool              `json:"published"`
	isOfficialArtist bool              `json:"published"`
	tags             []string          `json:"published"`
	channelIdType    int               `json:"published"`
	channelTabs      []string          `json:"published"`
	alertMessage     string            `json:"published"`
	channelLinks     ChannelLinks      `json:"published"`
}

type ChannelVideo struct {
	author          string         `json:"published"`
	authorId        string         `json:"published"`
	durationText    string         `json:"published"`
	lengthSeconds   int            `json:"published"`
	liveNow         bool           `json:"published"`
	premiere        bool           `json:"published"`
	premium         bool           `json:"published"`
	publishedText   string         `json:"published"`
	title           string         `json:"published"`
	Type            string         `json:"published"`
	videoId         string         `json:"published"`
	videoThumbnails []ChannelImage `json:"published"`
	viewCount       int            `json:"published"`
	viewCountText   string         `json:"published"`
}

type ChannelHomeItems struct {
	ShelfName string `json:"shelfName"`
	Type      string `json:"type"`
	Items     string `json:"items"`
}

type ChannelHome struct {
	FeaturedVideo ChannelVideo     `json:"featuredVideo"`
	Items         ChannelHomeItems `json:"items"`
}

type ChannelVideos struct {
	channelIdType int
	alertMessage  string
	items         []ChannelVideo
	continuation  string
}

type ChannelVideosContinuation struct {
	items        []ChannelVideo
	continuation string
}

type ChannelPlaylist struct {
	author            string
	authorId          string
	authorUrl         string
	playlistId        string
	playlistThumbnail string
	playlistUrl       string
	title             string
	Type              string
	videoCount        int
}

type ChannelPlaylists struct {
	channelIdType int
	alertMessage  string
	items         []ChannelPlaylist
	continuation  string
}

type ChannelPlaylistsContinuation struct {
	items        []ChannelPlaylist
	continuation string
}

type ChannelSearch struct {
	items        []ChannelVideo
	continuation string
}

type ChannelSearchContinuation struct {
	items        []ChannelVideo
	continuation string
}

type RelatedChannelsContinuation struct {
	items        []RelatedChannel
	continuation string
}

type ChannelCommunityPost struct {
	postText         string
	postId           string
	author           string
	authorThumbnails string
	publishedText    string
	voteCount        string
	postContent      string
}

type ChannelCommunityPosts struct {
	channelIdType int
	innerTubeApi  string
	items         []ChannelCommunityPost
	continuation  string
}

type ChannelCommunityPostsContinuation struct {
	innerTubeApi string
	items        []ChannelCommunityPost
	continuation string
}

type ChannelStats struct {
	joinedDate int
	viewCount  int
	location   string
}

type ChannelBasicInfo struct {
	authorId           string
	author             string
	authorUrl          string
	authorThumbnails   []AuthorThumbnail
	authorThumbnailUrl string
	authorVerified     bool
	subCount           int
	videoCount         int
	description        string
}

type SubscribedChannelsResponse struct {
	channels     []ChannelBasicInfo
	channelCount int
}
