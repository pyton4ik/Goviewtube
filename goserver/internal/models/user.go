package models

type SettingsDto struct {
	miniplayer                       bool
	chapters                         bool
	theme                            string
	sponsorblockEnabled              bool
	sponsorblockSegmentSponsor       string
	sponsorblockSegmentIntro         string
	sponsorblockSegmentOutro         string
	sponsorblockSegmentInteraction   string
	sponsorblockSegmentSelfpromo     string
	sponsorblockSegmentMusicOfftopic string
	sponsorblockSegmentPreview       string
	autoplay                         bool
	saveVideoHistory                 bool
	showHomeSubscriptions            bool
	alwaysLoopVideo                  bool
	autoplayNextVideo                bool
	audioModeDefault                 bool
	defaultVideoSpeed                int
	defaultVideoQuality              string
	defaultAudioQuality              string
	autoAdjustAudioQuality           bool
	autoAdjustVideoQuality           bool
	dashPlaybackEnabled              bool
}

type UserprofileDto struct {
	username     string
	profileImage string
	settings     SettingsDto
}
