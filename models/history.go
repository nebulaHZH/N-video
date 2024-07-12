package models

type History struct {
	Uid     int    `json:"uid"`     //用户id
	History string `json:"history"` //浏览历史
}
