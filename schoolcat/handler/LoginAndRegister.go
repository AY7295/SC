package handler

import (
	"SchoolCat/database"
	"SchoolCat/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EmailExist(email string) bool { //检查名字
	var user model.User
	DB := database.Link()
	res := DB.Where("email = ?", email).Take(&user)
	return res.RowsAffected != 0
}

func PasswordRight(pwd string, email string) bool { //检查密码
	var user model.User
	DB := database.Link()
	res := DB.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		log.Println(res.Error)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

func AdminExist(email string) bool {
	var admin model.Admin
	DB := database.Link()
	res := DB.Where("email = ?", email).Take(&admin)
	return res.RowsAffected != 0
}

func Login(c *gin.Context) { //登录
	var user,user0 model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(user.Password, user.Email)
	if !EmailExist(user.Email) || !PasswordRight(user.Password, user.Email) {
		c.AsciiJSON(400, gin.H{
			"msg": "用户名或者密码有误",
		})
		return
	}else if AdminExist(user.Email){
		c.AsciiJSON(200, gin.H{
			"msg": "欢迎使用,admin",
			"userid":user0.ID,//登陆时会传给前端一个ID，用户的每个操作都要返回一个ID，大写的
			"auth":"12345",//管理员地任何操作都要在header里面加上这个键值对
			"username":user0.Username,
			"iconsrc":user0.IconSrc,//这2个价值对在用户评论和发表分享时返回
		})
		return
	}
	DB := database.Link()
	res := DB.Where("email = ?", user.Email).Take(&user0)
	if res.Error != nil {
		log.Println(res.Error)
	}
	c.AsciiJSON(200, gin.H{
		"msg": "欢迎使用",
		"userid":user0.ID,//登陆时会传给前端一个ID，用户的每个操作都要返回一个ID，大写的

		"username":user0.Username,
		"iconsrc":user0.IconSrc,//这2个价值对在用户评论和发表分享时返回
	})

	//fmt.Println(EmailExist(user.Email), PasswordRight(user.Password, user.Email))
}

func Register(c *gin.Context) { //注册
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Panic(err)
	}
	if EmailExist(user.Email) {
		c.AsciiJSON(400, gin.H{
			"msg": "邮箱已被注册",
		})
		return
	} else {

		hash, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err1 != nil {
			fmt.Println(err1)
		}
		//fmt.Println(hash)
		encodePW := string(hash)
		//fmt.Println(encodePW)
		user.Password = encodePW
		DB := database.Link()
		//fmt.Println(user)
		err = DB.Create(&user).Error
		if err != nil {
			log.Println(err)
			return
		}
		c.AsciiJSON(200, gin.H{
			"msg": "注册成功",
		})
	}
}

func Info(c *gin.Context) {
	var user0, user model.User
	err := c.ShouldBind(&user0)
	if err != nil {
		log.Println(err)
		return
	}
	DB := database.Link()
	res := DB.Where("ID = ?", user0.ID).Take(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	user0.Email=user.Email
	user0.Password=user.Password
	DB.Save(&user0)
	c.AsciiJSON(200, gin.H{
		"msg": "操作成功",
		"username":user0.Username,
		"iconsrc":user0.IconSrc,//如果用户没上传头像就为空
	})
}