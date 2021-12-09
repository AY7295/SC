package handler

import (
	"SchoolCat/database"
	"SchoolCat/model"
	"fmt"
	"github.com/gin-gonic/gin"
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
	return pwd == user.Password
}

func AdminExist(email string) bool {
	var admin model.Admin
	DB := database.Link()
	res := DB.Where("email = ?", email).Take(&admin)
	return res.RowsAffected != 0
}

func Login(c *gin.Context) { //登录
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}
	if !EmailExist(user.Email) || !PasswordRight(user.Password, user.Email) {
		c.AsciiJSON(400, gin.H{
			"msg": "用户名或者密码有误",
		})
		return
	}else if AdminExist(user.Email){
		c.AsciiJSON(200, gin.H{
			"msg": "欢迎使用,admin",
		})
		return
	}
	c.AsciiJSON(200, gin.H{
		"msg": "欢迎使用",
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
		DB := database.Link()
		fmt.Println(user)
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

func AddInfo(c *gin.Context) {
	var user0, user model.User
	err := c.ShouldBind(&user0)
	if err != nil {
		log.Println(err)
		return
	}
	DB := database.Link()
	res := DB.Where("email = ?", user0.Email).Take(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	user.Username = user0.Username
	user.Gender = user0.Gender
	user.School = user0.School
	DB.Save(&user)
	c.AsciiJSON(200, gin.H{
		"msg": "添加成功",
	})
}
func UpdateInfo(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}
	DB := database.Link()
	//res := DB.Where("email = ?", user0.Email).Take(&user)
	//if res.Error != nil {
	//	log.Println(res.Error)
	//	return
	//}
	//if user0.Username!=user.Username && user0.Username!=""{user.Username = user0.Username}
	//if user0.Gender!=user.Gender && user0.Gender!=""{user.Gender = user0.Gender}
	//if user0.School!=user.School && user0.School!=""{user.School = user0.School}
	//fmt.Println(user0.Username,user0.Gender,user0.School,"	",user.Username,user.Gender,user.School)
	DB.Save(&user)
	c.AsciiJSON(200, gin.H{
		"msg": "更改成功",
	})
}