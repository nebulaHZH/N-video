package user

import (
	"N-video/models"
	"N-video/service/video"
)

// service: get author info
func GetAuthorInfo(uid int) models.AuthorImpl {
	author = models.AuthorImpl{Uid: uid}
	return author.Get_author_info()
}

// service: get author's folder list
func GetFolderList(uid int) []models.VideosImpl {
	author = models.AuthorImpl{Uid: uid}
	fids := author.Get_user_folder()
	var videos []models.VideosImpl
	for _, fid := range fids {
		videos = append(videos, video.GetFolderInfo(fid.Fid))
	}
	return videos
}
