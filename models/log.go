package models

type Log struct {
	Uid      int    `json:"uid"`      //用户id 传json字符串时用后面的id
	Password string `json:"password"` //密码
}

type LogInter interface {
	Valid(uid int, password string) Log
	UpdatePassword(uid int, password string)
	Create() Log
}

// dao: valid user's password
func (u Log) Valid(uid int, password string) Log {
	db.Take(&u, "Uid = ?", uid, "Password = ?", password)
	return u
}

// dao: update user's password
func (u Log) UpdatePassword(uid int, password string) {
	db.Model(&u).Where("Uid = ?", uid).Update("Password", password)
}

// dao: create user's uid——password
func (u Log) Create() Log {
	db.Create(&u)
	return u
}
