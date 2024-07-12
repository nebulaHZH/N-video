package user

import (
	user "N-video/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.PostForm("uid"))
	password := c.PostForm("password")
	// 调用model 方法操纵数据库
	person := user.User{Uid: uid, Password: password}
	p, err := person.GetUserInfo(uid)
	if err != nil {
		if p.Password == password {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "登录成功",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "密码错误",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
	}
}
func Rregister_handler(c *gin.Context) {
	//将post传来的json转为结构体
	person := user.User{}
	c.BindJSON(&person)
	person.AddUser()
}

func Modify_password_handler(c *gin.Context) {

}
