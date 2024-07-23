package models

type Followed struct {
	Uid      int `json:"uid"`      //用户id
	Followed int `json:"followed"` //关注的人
}

type FollowedInter interface {
	Get(fid string, uid int) []Followed
	Delete()
}

// dao : get all followed/follower by uid
func (f Followed) Get(t string, id int) []Followed {
	var followed_list []Followed
	db.Model(&f).Where(t+" = ?", id).Find(&followed_list)
	return followed_list
}

// dao: delete follow
func (f Followed) Delete() {
	db.Where("uid = ? AND followed = ?", f.Uid, f.Followed).Delete(&f)
}
