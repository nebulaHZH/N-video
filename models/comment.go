package models

type Comment struct {
	Cid        string `json:"cid"`        //评论id
	Vid        string `json:"vid"`        //视频id
	Sender     int    `json:"sender"`     //发送者id
	Receiver   int    `json:"receiver"`   //接收者id
	Content    string `json:"content"`    //评论内容
	CreateTime string `json:"createTime"` //评论时间
	Image      string `json:"image"`      // 评论图片
}

type CommentInter interface {
	GetComment() Comment
	SetComment()
	DeleteCommentBS(sender int, cid string)
	GetCommentsByVid(vid string) []Comment
	DeleteCommentByVid(vid string)
}

func (c Comment) GetComment() Comment {
	var comment Comment
	db.Model(&c).First(&comment)
	return comment
}
func (c Comment) GetCommentsByVid(vid string) []Comment {
	var comments []Comment
	db.Model(&c).Where("vid = ?", vid).Find(&comments)
	return comments
}
func (c Comment) SetComment() {
	db.Create(&c)
}
func (c Comment) DeleteCommentBS(sender int, cid string) {
	db.Where("sender = ? and cid = ?", sender, cid).Delete(&c)
}
func (c Comment) DeleteCommentByVid(vid string) {
	db.Where("vid = ?", vid).Delete(&c)
}
