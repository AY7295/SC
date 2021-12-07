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
	if user0.Username!=user.Username && user0.Username!=""{user.Username = user0.Username}
	if user0.Gender!=user.Gender && user0.Gender!=""{user.Gender = user0.Gender}
	if user0.School!=user.School && user0.School!=""{user.School = user0.School}
	fmt.Println(user0.Username,user0.Gender,user0.School,"	",user.Username,user.Gender,user.School)
	DB.Save(&user)
	c.AsciiJSON(200, gin.H{
		"msg": "更改成功",
	})
}

//func CompareStrings (a *string,b *string){
//	if a!=b {//传过来的的信息已经更改并且不是空的
//		a=b
//	}
//}

func NewShare(c *gin.Context){//每次产生一个share时应该生成一个moreshare
	//
	//
	//
	//
	var share model.Share
	err:=c.ShouldBind(&share)
	if err!=nil{
		fmt.Println(err);return
	}
	DB := database.Link()
	fmt.Println(share)
	err = DB.Create(&share).Error
	if err != nil {
		log.Println(err)
		return
	}

	var shareSrc []model.ShareImg
	err =c.ShouldBind(&shareSrc)//存地址

	var moreShare model.MoreShare
	err =c.ShouldBind(&moreShare)

	if err!=nil{
		fmt.Println(err);return
	}
	//......


}

func ViewShareText(c *gin.Context){//传share的文字
	id := c.Query("shareid")
	var share model.Share
	DB := database.Link()
	res := DB.Where("id = ?", id).Take(&share)
	if res.Error!=nil{fmt.Println(res.Error);return}


}
func ViewShareSrc (c *gin.Context){//传share的图片
	id := c.Query("shareid")
	var src []model.ShareImg
	DB := database.Link()
	DB.Find(&src).Where("shareid= ?",id)
	c.AsciiJSON(200,gin.H{
		"src":src,
	})
}


func MoreShare (c *gin.Context){//首页刷新给的share
	var cat []model.MoreShare
	DB := database.Link()
	DB.Find(&cat).Limit(10).Offset(10)
	//fmt.Println(cat)
	c.AsciiJSON(200,gin.H{
		"cats":cat,
	})
}
func Search(c *gin.Context){//搜索
	keywords:=c.Query("keywords")
	DB:=database.Link()
	var cat []model.MoreShare
	DB.Where("title LIKE ?","%"+keywords+"%").Find(&cat).Limit(10).Offset(10)
	c.AsciiJSON(200,gin.H{
		"cats":cat,
	})
}

func UpdateTip(c *gin.Context){

}
func AddTip (c *gin.Context){

}

func DeleteTip(c *gin.Context){

}