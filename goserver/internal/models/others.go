package models

type Chapter struct {
	title     string `json:"title"`
	startTime int    `json:"startTime"`
}

type Dislike struct {
	Id          string `json:"id"`
	DateCreated string `json:"startTime"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Rating      int    `json:"rating"`
	ViewCount   int    `json:"viewCount"`
	Deleted     bool   `json:"deleted"`
}

type FilterValue struct {
	active      bool   `json:"startTime"`
	description string `json:"startTime"`
	name        string `json:"startTime"`
	url         string `json:"startTime"`
}

type SearchFilter struct {
	filterType   string        `json:"startTime"`
	filterValues []FilterValue `json:"startTime"`
}

type SubscriptionStatus struct {
	channelId    string `json:"startTime"`
	isSubscribed bool   `json:"startTime"`
}

type Registration struct {
	username        string `json:"startTime"`
	password        string `json:"startTime"`
	captchaToken    string `json:"startTime"`
	captchaSolution string `json:"startTime"`
}

type Captcha struct {
	token        string `json:"startTime"`
	captchaImage string `json:"startTime"`
}
