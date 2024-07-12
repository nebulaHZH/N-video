package models

import (
	mysql "N-video/utils"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Uid         int    `json:"uid"`        //用户id
	Username    string `json:"username"`   //用户名
	Sex         string `json:"sex"`        //性别
	Background  string `json:"background"` //背景
	Avatar      string `json:"avatar"`     //头像
	Intrduction string `json:"moto"`       //简介
	Truename    string `json:"truename"`   //真实名
	Password    string `json:"password"`   //密码
	IsVIP       bool   `json:"isVIP"`      //是否为vip
}
type Person interface {
	GetUserInfo(uid int) (User, error)             //获取用户信息
	UpdateUserInfo(user User) bool                 //更新用户信息
	DeleteUser(uid int) bool                       //删除用户
	AddUser(user User) bool                        //添加用户
	GetAuthors(uid int) (Authors, error)           //获取作者信息
	GetUserVideoFolder(fid int) ([]Videos, error)  //获取用户视频文件夹
	GetUserFollowers(uid int) ([]User, error)      //获取用户粉丝
	GetUserFollowed(uid int) ([]User, error)       //获取用户关注
	CreateVideoFolder(uid int, folder Videos) bool //用户创建文件夹
	GetHistory(uid int) ([]Videos, error)          //获取用户历史记录
}

func init() {
	db = mysql.GetDB()
}

// 在定义模型时指定表名为 "user"
func (User) TableName() string {
	return "user"
}
func (u *User) GetUserInfo(uid int) (User, error) {
	db.First(&u, uid)
	return *u, nil
}
func (u *User) UpdateUserInfo(user User) bool {
	return true
}
func (u *User) DeleteUser(uid int) bool {
	return true
}
func (u *User) AddUser() bool {
	db.Create(&u)
	return true
}
func (u *User) GetAuthors(uid int) (Authors, error) {
	return Authors{}, nil
}
func (u *User) GetUserVideoFolder(fid int) ([]Videos, error) {
	return []Videos{}, nil
}
func (u *User) GetUserFollowers(uid int) ([]User, error) {
	return []User{}, nil
}
func (u *User) GetUserFollowed(uid int) ([]User, error) {
	return []User{}, nil
}
func (u *User) CreateVideoFolder(uid int, folder Videos) bool {
	return true
}
func (u *User) GetHistory(uid int) ([]Videos, error) {
	return []Videos{}, nil
}
