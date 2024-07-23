package user

import (
	"N-video/models"
)

// service: get author's folder list
func GetFolderList(uid int) []models.Author {
	author = models.Author{Uid: uid}
	return author.Get_folder()
}

// user create a folder , need uid and folder name
func CreateFolder(uid int, foldername string) bool {
	author = models.Author{Uid: uid, Fname: foldername}
	author.Create_folder()

	return author.Get_folder()[0].Fid != ""
}
