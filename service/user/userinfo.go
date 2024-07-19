package user

import (
	"N-video/models"
)

func GetUserInfo(uid int) models.User {
	u, _ := person.GetUserInfo(uid)
	return u
}

func UpdateUserInfo(user models.User) models.User {
	user.UpdateUserInfo()
	// if update sucess , we can find it by GetuserInfo function
	return GetUserInfo(user.Uid)
}

func DeleteUserInfo(uid int) models.User {
	person.DeleteUser(uid)
	return GetUserInfo(uid)
}

// service : get followed by user's id
func GetFollowed(uid int) []models.User {
	f_list := followed.Get("uid", uid)
	var user_list []models.User
	for item := range f_list {
		u, _ := person.GetUserInfo(f_list[item].Followed)
		user_list = append(user_list, u)
	}
	return user_list
}

// service : get follower by user's id
func GetFollower(uid int) []models.User {
	follower_list := followed.Get("fid", uid)
	var followers []models.User
	for item := range follower_list {
		u, _ := person.GetUserInfo(follower_list[item].Uid)
		followers = append(followers, u)
	}
	return followers
}

// service : cancel followed
func CancelFollow(uid, fid int) bool {
	followed = models.FollowedImpl{Uid: uid, Followed: fid}
	followed.Delete()
	// judge if success to cancel
	res := followed.Get("uid", uid)
	return len(res) == 0
}
