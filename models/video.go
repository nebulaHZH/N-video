package models

// video model
type Video struct {
	Vid           string `json:"vid"`           //视频id
	Title         string `json:"title"`         //标题
	Uid           int    `json:"uid"`           //作者id
	Createtime    string `json:"createTime"`    //创建时间
	Playurl       string `json:"playUrl"`       //视频链接
	Coverurl      string `json:"coverUrl"`      //封面图片
	Favoritecount int    `json:"favoriteCount"` //点赞数
	Commentcount  int    `json:"commentCount"`  //评论数
	Viewcount     int64  `json:"viewCount"`     //浏览数
	Isneedvip     bool   `json:"isNeedVip"`     //是否需要vip
}

// relation model
type Relation struct {
	Fid string `json:"fid"` //视频集id
	Vid string `json:"vid"` //视频id
}

type VideoInter interface {
	GetVideo() Video
	UploadVideo() bool
	DeleteVideo()
}
type RelationInter interface {
	AddIntoVideoFolder(fid string, vid string)
	RemoveFromVideoFolder(fid string, vid string)
}

func (v Video) UploadVideo() bool {
	return true
}
func (v Video) GetVideo() Video {
	db.First(&v, v.Vid)
	return v
}
func (v Video) DeleteVideo() {
	db.Where("vid = ?", v.Vid).Delete(&v)
}

func (v Relation) AddIntoVideoFolder(fid string, vid string) {
	db.Create(&v)
}
func (v Relation) RemoveFromVideoFolder(fid string, vid string) {
	db.Delete(&v)
}
