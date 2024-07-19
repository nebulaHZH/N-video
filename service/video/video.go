package video

import (
	"N-video/models"
)

func GetVideoInfo(vid string) models.VideoImpl {
	video = models.VideoImpl{Vid: vid}
	return video.GetVideo()
}
func DeleteVideo(vid string) bool {
	video = models.VideoImpl{Vid: vid}
	video.DeleteVideo()
	return GetVideoInfo(vid).Vid != ""
}
