package util

import (
	"SchoolCat/model"
	"github.com/gin-gonic/gin"
)

func UserNotExist (c *gin.Context)  {
	c.AsciiJSON(400, gin.H{
		"msg": "不存在该用户",
	})
}

func PasswordWrong (c *gin.Context )  {
	c.AsciiJSON(400, gin.H{
		"msg": "密码错误",
	})
}

func AdminLogin(c *gin.Context,user0 model.User,token string)  {
	c.AsciiJSON(200, gin.H{
		"msg": "欢迎使用,admin",
		"userid":user0.ID,//登陆时会传给前端一个ID，用户的每个操作都要返回一个ID，大写的
		"auth":"12345",//管理员地任何操作都要在header里面加上这个键值对
		"username":user0.Username,
		"icon":user0.IconSrc,//这2个价值对在用户评论和发表分享时返回
		"token":token,
	})
}

func UserLogin(c *gin.Context,user0 model.User,token string)  {
	c.AsciiJSON(200, gin.H{
		"msg": "欢迎使用",
		"userid":user0.ID,//登陆时会传给前端一个ID，用户的每个操作都要返回一个ID，大写的
		"username":user0.Username,
		"icon":user0.IconSrc,//这2个价值对在用户评论和发表分享时返回
		"token":token,
	})
}

func InfoLost(c *gin.Context)  {
	c.AsciiJSON(400,gin.H{
		"msg":"信息不完整",
	})
}

func EmailRegistered(c *gin.Context)  {
	c.AsciiJSON(400, gin.H{
		"msg": "邮箱已被注册",
	})
}

func RegisterSucceed(c *gin.Context, userId uint)  {
	c.AsciiJSON(200, gin.H{
		"msg":     "注册成功",
		"user_id": userId,
	})
}

func UpdateInfo(c *gin.Context,user0 model.User)  {
	c.AsciiJSON(200, gin.H{
		"msg": "操作成功",
		"username":user0.Username,
		"icon":user0.IconSrc,//如果用户没上传头像就为空
	})
}

