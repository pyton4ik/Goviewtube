package models

type CommentDto struct {
	authorThumbnails AuthorThumbnail
	author           string
	authorId         string
	publishedText    string
	isEdited         bool
	likeCount        int
	creatorHeart     bool
	replyToken       string
	replyCount       int
	content          string
}

type CommentsResponseDto struct {
	comments     []CommentDto
	continuation string
}
