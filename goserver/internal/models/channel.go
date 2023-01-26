package models

type ChannelImageDto struct {
	url    string
	width  int
	height int
}

type RelatedChannelDto struct {
	channelName     string
	channelId       string
	channelUrl      string
	thumbnail       []ChannelImageDto
	videoCount      int
	subscriberText  string
	subscriberCount int
	verified        bool
	officialArtist  bool
	officialArist   bool
}

type ChannelLinkDto struct {
	url   string
	icon  string
	title string
}

type RelatedChannels struct {
	items        []RelatedChannelDto
	continuation string
}

type ChannelLinks struct {
	primaryLinks   []ChannelLinkDto
	secondaryLinks []ChannelLinkDto
}

type ChannelInfoDto struct {
	author           string
	authorId         string
	authorUrl        string
	authorBanners    []ChannelImageDto
	authorThumbnails []ChannelImageDto
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

type ChannelVideoDto struct {
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
	videoThumbnails []ChannelImageDto
	viewCount       int
	viewCountText   string
}

type ChannelHomeDtoItems struct {
	shelfName string
	Type      string
	items     string
}

type ChannelHomeDto struct {
	featuredVideo ChannelVideoDto
	items         ChannelHomeDtoItems
}

type ChannelVideosDto struct {
	channelIdType int
	alertMessage  string
	items         []ChannelVideoDto
	continuation  string
}

type ChannelVideosContinuationDto struct {
	items        []ChannelVideoDto
	continuation string
}

type ChannelPlaylistDto struct {
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

type ChannelPlaylistsDto struct {
	channelIdType int
	alertMessage  string
	items         []ChannelPlaylistDto
	continuation  string
}

type ChannelPlaylistsContinuationDto struct {
	items        []ChannelPlaylistDto
	continuation string
}

type ChannelSearchDto struct {
	items        []ChannelVideoDto
	continuation string
}

type ChannelSearchContinuationDto struct {
	items        []ChannelVideoDto
	continuation string
}

type RelatedChannelsContinuationDto struct {
	items        []RelatedChannelDto
	continuation string
}

type ChannelCommunityPostDto struct {
	postText         string
	postId           string
	author           string
	authorThumbnails string
	publishedText    string
	voteCount        string
	postContent      string
}

type ChannelCommunityPostsDto struct {
	channelIdType int
	innerTubeApi  string
	items         []ChannelCommunityPostDto
	continuation  string
}

type ChannelCommunityPostsContinuationDto struct {
	innerTubeApi string
	items        []ChannelCommunityPostDto
	continuation string
}

type ChannelStatsDto struct {
	joinedDate int
	viewCount  int
	location   string
}

type ChannelBasicInfoDto struct {
	authorId           string
	author             string
	authorUrl          string
	authorThumbnails   []AuthorThumbnailDto
	authorThumbnailUrl string
	authorVerified     bool
	subCount           int
	videoCount         int
	description        string
}

type SubscribedChannelsResponseDto struct {
	channels     []ChannelBasicInfoDto
	channelCount int
}
