package models

type FollowedImpl struct {
	Uid      int `json:"uid"`      //用户id
	Followed int `json:"followed"` //关注的人
}

type Followed interface {
	Get(fid string, uid int) []FollowedImpl
	Delete()
}

// dao : get all followed/follower by uid
func (f FollowedImpl) Get(t string, id int) []FollowedImpl {
	var followed_list []FollowedImpl
	db.Model(&f).Where(t+" = ?", id).Find(&followed_list)
	return followed_list
}

// dao: delete follow
func (f FollowedImpl) Delete() {
	db.Where("uid = ? AND followed = ?", f.Uid, f.Followed).Delete(&f)
}
