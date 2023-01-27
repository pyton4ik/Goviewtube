package models

type ChannelImage struct {
	url    string
	width  int
	height int
}

type RelatedChannel struct {
	channelName     string
	channelId       string
	channelUrl      string
	thumbnail       []ChannelImage
	videoCount      int
	subscriberText  string
	subscriberCount int
	verified        bool
	officialArtist  bool
	officialArist   bool
}

type ChannelLink struct {
	url   string
	icon  string
	title string
}

type RelatedChannels struct {
	items        []RelatedChannel
	continuation string
}

type ChannelLinks struct {
	primaryLinks   []ChannelLink
	secondaryLinks []ChannelLink
}

type ChannelInfo struct {
	author           string
	authorId         string
	authorUrl        string
	authorBanners    []ChannelImage
	authorThumbnails []ChannelImage
	subscriberText   string
	subscriberCount  int
	description      string
	isFamilyFriendly bool
	relatedChannels  RelatedChannels
	allowedRegions   []string
	isVerified       bool
	isOfficialArtist bool
	tags             []string
	channelIdType    int
	channelTabs      []string
	alertMessage     string
	channelLinks     ChannelLinks
}

type ChannelVideo struct {
	author          string
	authorId        string
	durationText    string
	lengthSeconds   int
	liveNow         bool
	premiere        bool
	premium         bool
	publishedText   string
	title           string
	Type            string
	videoId         string
	videoThumbnails []ChannelImage
	viewCount       int
	viewCountText   string
}

type ChannelHomeItems struct {
	shelfName string
	Type      string
	items     string
}

type ChannelHome struct {
	featuredVideo ChannelVideo
	items         ChannelHomeItems
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
