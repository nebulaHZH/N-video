package models

type AuthorImpl struct {
	Uid   int    `json:"uid"`   //用户id
	Fid   string `json:"works"` //视频集
	Fname string `json:"fname"` //视频集名
}

type Author interface {
	Get_user_folder() []AuthorImpl
	Get_author_info() AuthorImpl
}

// this function is used to get user's works[params: uid]
func (a AuthorImpl) Get_user_folder() []AuthorImpl {
	var fid_list []AuthorImpl
	db.Where("uid = ?", a.Uid).Find(&fid_list)
	return fid_list
}

// dao: get author information
func (a AuthorImpl) Get_author_info() AuthorImpl {
	db.Where("uid = ?", a.Uid).First(&a)
	return a
}
