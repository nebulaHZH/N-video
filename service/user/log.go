package user

import (
	"N-video/models"
)

// service: valid user's login
func Login(uid int, password string) bool {
	res := log.Valid(uid, password)
	return res.Uid != 0
}

// service: register account
func Register(user models.User, password string) bool {
	logImpl := models.LogImpl{Uid: user.Uid, Password: password}
	var log models.Log = logImpl
	res, l := user.AddUser(), log.Create()
	if res.Uid != 0 && l.Uid != 0 {
		return true
	}
	return false
}

// service: modify user's password
func ModifyPsd(uid int, password string) bool {
	log.UpdatePassword(uid, password)
	person := log.Valid(uid, password)
	return person.Uid != 0
}
