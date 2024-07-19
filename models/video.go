package models

type VideoImpl struct {
	Vid           string    `json:"vid"`           //视频id
	Title         string    `json:"title"`         //标题
	Uid           int       `json:"uid"`           //作者id
	CreateTime    string    `json:"createTime"`    //创建时间
	PlayUrl       string    `json:"playUrl"`       //视频链接
	CoverUrl      string    `json:"coverUrl"`      //封面图片
	FavoriteCount int       `json:"favoriteCount"` //点赞数
	CommentCount  int       `json:"commentCount"`  //评论数
	ViewCount     int64     `json:"viewCount"`     //浏览数
	IsNeedVIP     bool      `json:"isNeedVip"`     //是否需要vip
	Comments      []Comment `json:"comments"`      //评论实体
}
type VideosImpl struct {
	Fid   string `json:"fid"`   //视频集id
	Fname string `json:"fname"` //视频集名称
	Uid   int    `json:"uid"`   //用户id
	Vid   string `json:"vid"`   //视频id列表
}

type Video interface {
	GetVideo() VideoImpl
	UploadVideo() bool
	DeleteVideo()
}
type Folder interface {
	GetVideoFolder() VideosImpl
	AddIntoVideoFolder(fid string, vid string) error
	RemoveFromVideoFolder(fid string, vid string) error
}

func (v VideoImpl) UploadVideo() bool {
	return true
}
func (v VideoImpl) GetVideo() VideoImpl {
	db.First(&v, v.Vid)
	return v
}
func (v VideoImpl) DeleteVideo() {
	db.Where("vid = ?", v.Vid).Delete(&v)
}

func (v VideosImpl) GetVideoFolder() VideosImpl {
	db.Where("fid = ? AND uid = ?", v.Fid, v.Uid).First(&v)
	return v
}
func (v VideosImpl) AddIntoVideoFolder(fid string, vid string) error {
	return nil
}
func (v VideosImpl) RemoveFromVideoFolder(fid string, vid string) error {
	return nil
}
