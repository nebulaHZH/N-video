package models

type Author struct {
	Uid   int    `json:"uid"`   //用户id
	Fid   string `json:"works"` //视频集
	Fname string `json:"fname"` //视频集名
}
type AuthorInter interface {
	Get_folder() []Author
	Create_folder()
}

// this function is used to get user's works[params: uid]
func (a Author) Get_folder() []Author {
	var fid_list []Author
	db.Where("uid = ?", a.Uid).Find(&fid_list)
	return fid_list
}

// model : create folder [params: uid, works, fname]
func (a Author) Create_folder() {
	db.Create(&a)
}
