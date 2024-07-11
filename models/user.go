package models

import "github.com/gin-gonic/gin"

type User struct {
	Uid            int      `json:"uid"`            //用户id
	Username       string   `json:"username"`       //用户名
	Sex            string   `json:"sex"`            //性别
	Background     string   `json:"background"`     //背景
	Avatar         string   `json:"avatar"`         //头像
	Intrduction    string   `json:"moto"`           //简介
	Truename       string   `json:"truename"`       //真实名
	Password       string   `json:"password"`       //密码
	IsVIP          bool     `json:"isVIP"`          //是否为vip
	History        []string `json:"history"`        //浏览历史
	Followed       []string `json:"followed"`       //关注的人
	Authors        Authors  `json:"authors"`        //创作者身份
	FavoriteVideos []string `json:"favoriteVideos"` //喜欢的视频
}

type Authors struct {
	Uid         int      `json:"uid"`       //用户id
	VideoFolder []Videos `json:"works"`     //视频集
	Followers   []string `json:"followers"` //关注者
}

type Person interface {
	GetUserInfo(c *gin.Context) (User, error)      //获取用户信息
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

func (u *User) GetUserInfo(uid int) (User, error) {
	return *u, nil
}
func (u *User) UpdateUserInfo(user User) bool {
	return true
}
func (u *User) DeleteUser(uid int) bool {
	return true
}
func (u *User) AddUser(user User) bool {
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
