package user

import (
	"N-video/models"
	"N-video/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type fullUser struct {
	models.User
	Password string `json:"password"`
}

// 用户登录处理函数，所需参数：账号，密码，目前还没加token等
func Login_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.PostForm("uid"))
	password := c.PostForm("password")
	c.JSON(200, user.Login(uid, password))
}

// user register function , you can add some middleware in it
func Rregister_handler(c *gin.Context) {
	//将post传来的json转为结构Background体
	fuser := fullUser{}
	//获取post的json参数
	err := c.BindJSON(&fuser)
	if err != nil {
		c.JSON(200, gin.H{"msg": "参数错误"})
		return
	}
	/*middleware's place*/
	res := user.Register(fuser.User, fuser.Password)
	c.JSON(200, gin.H{"msg": res})
}

// user modify password function , you can add some middleware in it
func Modify_password_handler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.PostForm("uid"))
	new_password := c.PostForm("password")
	res := user.ModifyPsd(uid, new_password)
	c.JSON(200, gin.H{"msg": res})
}
