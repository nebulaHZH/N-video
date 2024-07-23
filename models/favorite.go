package models

type Favorite struct {
	Vid string `json:"vid"` //视频id
	Uid int    `json:"uid"` //用户id
}

type FavoriteInter interface {
	Like()
	Dislike()
	GetLike(vid string) Favorite
}

func (f Favorite) Like() {
	db.Create(&f)
}

func (f Favorite) Dislike() {
	db.Delete(&f)
}

func (f Favorite) GetLike(vid string) Favorite {
	var res Favorite
	db.First(&res, "vid = ?", vid)
	return res
}
