package models

type PlaylistImageDto struct {
	url    string
	width  int
	height int
}

type PlaylistItemDtoAuthor struct {
	name      string
	url       string
	channelID string
}
type PlaylistItemDto struct {
	title         string
	index         int
	id            string
	shortUrl      string
	url           string
	author        PlaylistItemDtoAuthor
	thumbnails    []PlaylistImageDto
	bestThumbnail PlaylistImageDto
	isLive        bool
	duration      string
	durationSec   int
}

type PlaylistResultDtoAuthor struct {
	name       string
	url        string
	avatars    []PlaylistImageDto
	bestAvatar PlaylistImageDto
	channelID  string
}
type PlaylistResultDto struct {
	id                 string
	url                string
	title              string
	estimatedItemCount int
	views              int
	thumbnails         []PlaylistImageDto
	bestThumbnail      PlaylistImageDto
	lastUpdated        string
	description        string
	visibility         string
	author             PlaylistResultDtoAuthor
	items              []PlaylistItemDto
	continuation       string
}
