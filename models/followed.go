package models

type Followed struct {
	Uid      int    `json:"uid"`      //用户id
	Followed string `json:"followed"` //关注的人
}
