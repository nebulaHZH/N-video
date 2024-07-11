package models

type Comment struct {
	Cid        string `json:"cid"`        //评论id
	Sender     int    `json:"sender"`     //发送者id
	Receiver   int    `json:"receiver"`   //接收者id
	Content    string `json:"content"`    //评论内容
	CreateTime string `json:"createTime"` //评论时间
	Image      string `json:"image"`      // 评论图片
}

type Review interface {
	GetComment(cid string) (Comment, error)
	SetComment(comment Comment) bool
	DeleteComment(cid string) bool
}

func (c *Comment) GetComment(cid string) (Comment, error) {
	return *c, nil
}
func (c *Comment) SetComment(comment Comment) bool {
	*c = comment
	return true
}
func (c *Comment) DeleteComment(cid string) bool {
	return true
}
