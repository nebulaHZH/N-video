package models

import (
	mysql "N-video/utils"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	Uid          int    `json:"uid"`          //用户id 传json字符串时用后面的id
	Username     string `json:"username"`     //用户名
	Sex          string `json:"sex"`          //性别
	Background   string `json:"background"`   //背景
	Avatar       string `json:"avatar"`       //头像
	Introduction string `json:"introduction"` //简介
	Truename     string `json:"truename"`     //真实名
	Isvip        bool   `json:"isVIP"`        //是否为vip
}
type Person interface {
	GetUserInfo(uid int) (User, error)                 //获取用户信息
	UpdateUserInfo(user User) bool                     //更新用户信息
	DeleteUser(uid int) bool                           //删除用户
	AddUser(user User) bool                            //添加用户
	GetAuthors(uid int) (AuthorImpl, error)            //获取作者信息
	GetUserVideoFolder(fid int) ([]VideosImpl, error)  //获取用户视频文件夹
	GetUserFollowers(uid int) ([]User, error)          //获取用户粉丝
	GetUserFollowed(uid int) ([]User, error)           //获取用户关注
	CreateVideoFolder(uid int, folder VideosImpl) bool //用户创建文件夹
	GetHistory(uid int) ([]VideosImpl, error)          //获取用户历史记录
}

func init() {
	db = mysql.GetDB()
}

// 在定义模型时指定表名为 "user"
func (User) TableName() string {
	return "user"
}

// 获取用户信息，参数：uid
func (u *User) GetUserInfo(uid int) (User, error) {
	db.First(&u, uid)
	return *u, nil
}

// 更新用户信息
func (u *User) UpdateUserInfo() {
	db.Model(&u).Updates(u)
}

// delete user by uid
func (u *User) DeleteUser(uid int) {
	db.Model(&u).Delete(&u, uid)
}

// 添加用户
func (u *User) AddUser() User {
	db.Create(&u)
	return *u
}

func (u *User) GetAuthors(uid int) (AuthorImpl, error) {
	return AuthorImpl{}, nil
}
func (u *User) GetUserVideoFolder(fid int) ([]VideosImpl, error) {
	return []VideosImpl{}, nil
}
func (u *User) GetUserFollowers(uid int) ([]User, error) {
	return []User{}, nil
}
func (u *User) CreateVideoFolder(uid int, folder VideosImpl) bool {
	return true
}
func (u *User) GetHistory(uid int) ([]VideosImpl, error) {
	return []VideosImpl{}, nil
}
