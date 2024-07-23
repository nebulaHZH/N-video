package video

import (
	"N-video/models"
)

func GetVideoInfo(vid string, uid int) models.Video {
	video = models.Video{Vid: vid}
	var h models.History = models.History{Uid: uid, Vid: vid}
	h.Add()
	return video.GetVideo()
}
func DeleteVideo(vid string) bool {
	video = models.Video{Vid: vid}
	video.DeleteVideo()
	return video.GetVideo().Vid != ""
}
