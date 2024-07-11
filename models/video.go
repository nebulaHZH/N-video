package models

type Video struct {
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
type Videos struct {
	Fid    string  `json:"fid"`    //视频集id
	Uid    int     `json:"uid"`    //用户id
	Videos []Video `json:"videos"` //视频列表
}

type VideoInter interface {
	GetVideoInfo(vid string) (Video, error)
	GetComments(vid []string) ([]Comment, error)
}
type VideosFolder interface {
	GetVideoFolder(fid string) (Videos, error)
	AddIntoVideoFolder(fid string, vid string) error
	RemoveFromVideoFolder(fid string, vid string) error
}

func (v *Video) GetVideoInfo(vid string) (Video, error) {
	return *v, nil
}
func (v *Videos) GetVideoFolder(fid string) (Videos, error) {
	return *v, nil
}
func (v *Videos) AddIntoVideoFolder(fid string, vid string) error {
	return nil
}
func (v *Videos) RemoveFromVideoFolder(fid string, vid string) error {
	return nil
}
func (v *Video) GetComments(vid []string) ([]Comment, error) {
	return v.Comments, nil
}
