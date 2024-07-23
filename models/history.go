package models

type History struct {
	Uid int    `json:"uid"` //用户id
	Vid string `json:"vid"` //视频id
}

type HistoryInter interface {
	Get() []History
	Delete()
	Add()
}

func (h History) Get() []History {
	var h_list []History
	db.Where("uid = ?", h.Uid).Find(&h_list)
	return h_list
}

func (h History) Delete() {
	db.Where("uid = ? and vid = ?", h.Uid, h.Vid).Delete(&h)
}

func (h History) Add() {
	db.Create(&h)
}
