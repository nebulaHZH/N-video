package models

type LogImpl struct {
	Uid      int    `json:"uid"`      //用户id 传json字符串时用后面的id
	Password string `json:"password"` //密码
}

type Log interface {
	Valid(uid int, password string) LogImpl
	UpdatePassword(uid int, password string)
	Create() LogImpl
}

// dao: valid user's password
func (u LogImpl) Valid(uid int, password string) LogImpl {
	db.Take(&u, "Uid = ?", uid, "Password = ?", password)
	return u
}

// dao: update user's password
func (u LogImpl) UpdatePassword(uid int, password string) {
	db.Model(&u).Where("Uid = ?", uid).Update("Password", password)
}

// dao: create user's uid——password
func (u LogImpl) Create() LogImpl {
	db.Create(&u)
	return u
}
