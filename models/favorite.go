package models

type FavoriteImpl struct {
	Vid string `json:"vid"` //视频id
	Uid int    `json:"uid"` //用户id
}

type Favorite interface {
	Like(vid string, uid int)
	Dislike(vid string, uid int)
}

func (f FavoriteImpl) Like(vid string, uid int) {
	db.Create(&f)
}

func (f FavoriteImpl) Dislike(vid string, uid int) {
	db.Where("vid = ? AND uid = ?", vid, uid).Delete(&f)
}
