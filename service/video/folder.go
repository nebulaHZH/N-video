package video

import "N-video/models"

func GetFolderInfo(fid string) models.VideosImpl {
	folder = models.VideosImpl{Fid: fid}
	return folder.GetVideoFolder()
}
