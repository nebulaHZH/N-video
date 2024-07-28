package video

import (
	"N-video/models"
)

// service : add comment to video
func AddComment(c models.Comment) bool {
	comment = c
	comment.SetComment()
	return comment.GetComment().Cid != ""
}

func GetCommentList(videoId string) []models.Comment {
	comment = models.Comment{Vid: videoId}
	return comment.GetCommentsByVid(videoId)
}

func DeleteCommentBySender(sender int, cid string) bool {
	comment = models.Comment{Sender: sender, Cid: cid}
	comment.DeleteCommentBS(sender, cid)
	return comment.GetComment().Vid != ""
}

func DeleteCommentByVid(videoId string) bool {
	comment = models.Comment{Vid: videoId}
	comment.DeleteCommentByVid(videoId)
	return comment.GetComment().Vid != ""
}
