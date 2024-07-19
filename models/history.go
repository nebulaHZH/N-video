package models

type HistoryImpl struct {
	Uid int    `json:"uid"` //用户id
	Vid string `json:"vid"` //视频id
}

type History interface {
	Get() []HistoryImpl
	Delete()
}

func (h HistoryImpl) Get() []HistoryImpl {
	var h_list []HistoryImpl
	db.Where("uid = ?", h.Uid).Find(&h_list)
	return h_list
}

func (h HistoryImpl) Delete() {
	db.Where("uid = ? and vid = ?", h.Uid, h.Vid).Delete(&h)
}
