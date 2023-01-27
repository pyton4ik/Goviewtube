package models

type Chapter struct {
	title     string
	startTime int
}

type DislikeDto struct {
	id          string
	dateCreated string
	likes       int
	dislikes    int
	rating      int
	viewCount   int
	deleted     bool
}

type FilterValueDto struct {
	active      bool
	description string
	name        string
	url         string
}

type SearchFilterDto struct {
	filterType   string
	filterValues []FilterValueDto
}

type SubscriptionStatusDto struct {
	channelId    string
	isSubscribed bool
}

type RegistrationDto struct {
	username        string
	password        string
	captchaToken    string
	captchaSolution string
}

type CaptchaDto struct {
	token        string
	captchaImage string
}
