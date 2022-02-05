package service

import (
	"SchoolCat/database"
	"SchoolCat/midware"
	"SchoolCat/model"
	response "SchoolCat/util/responser"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var DB = database.DB

func EmailExist(email string) bool { //检查名字
	var user model.User

	res := DB.Where("email = ?", email).Take(&user)
	return res.RowsAffected != 0
}

func PasswordRight(pwd string, email string) bool { //检查密码
	var user model.User

	res := DB.Where("email = ?", email).Take(&user)
	if res.Error != nil {
		log.Println(res.Error)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	return err == nil
}

func AdminExist(email string) bool {
	var admin model.Admin

	res := DB.Where("email = ?", email).Take(&admin)
	return res.RowsAffected != 0
}

func Login(c *gin.Context) { //登录
	var user, user0 model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return
	}

	res := DB.Where("email = ?", user.Email).Take(&user0)

	if res.Error != nil {
		log.Println(res.Error)
	}
	if res.RowsAffected == 0 {
		response.UserNotExist(c)
		return
	}

	//fmt.Println(user.Password, user.Email)
	if !PasswordRight(user.Password, user.Email) {

		response.PasswordWrong(c)
		return
	}

	token := midware.GenerateToken(user.Email)

	if AdminExist(user.Email) {

		response.AdminLogin(c, user0, token)
		return
	}

	response.UserLogin(c, user0, token)

	//fmt.Println(EmailExist(user.Email), PasswordRight(user.Password, user.Email))
}

func Register(c *gin.Context) { //注册
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
	}

	if user.Password == "" || user.Email == "" {
		response.InfoLost(c)
		return
	}

	if EmailExist(user.Email) {
		response.EmailRegistered(c)
		return
	}

	hash, err1 := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Println(hash)
	encodePW := string(hash)
	//fmt.Println(encodePW)
	user.Password = encodePW

	//fmt.Println(user)
	err = DB.Create(&user).Error
	if err != nil {
		log.Println(err)
		return
	}

	response.RegisterSucceed(c, user.ID)

}

func Info(c *gin.Context) {
	var user0, user model.User
	err := c.ShouldBind(&user0)
	if err != nil {
		log.Println(err)
		return
	}

	res := DB.Where("ID = ?", user0.ID).Take(&user)
	if res.Error != nil {
		log.Println(res.Error)
		return
	}
	user0.Email = user.Email
	user0.Password = user.Password
	DB.Save(&user0)
	response.UpdateInfo(c, user0)
}
