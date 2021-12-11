package handler

import (
	"SchoolCat/database"
	"SchoolCat/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)



func Tip(c *gin.Context){
	var tip model.Tip
	err:=c.ShouldBind(&tip)
	if err != nil {
		fmt.Println(err); return
	}
	DB := database.Link()
	//fmt.Println(tip)
	err = DB.Create(&tip).Error
	if err != nil {
		log.Println(err)
	}
	c.AsciiJSON(200,gin.H{
		"msg":"操作成功",
		"CardID":tip.ID,
	})
}

func DeleteTip(c *gin.Context){
	DB := database.Link()
	var tip model.Tip
	tipid := c.GetHeader("tip_id")

	var user model.User
	uid,_ := strconv.Atoi(c.GetHeader("user_id"))//执行删除操作者的id
	res := DB.Where("id = ?",uid).Take(&user)
	if res.Error != nil{fmt.Println(res.Error);return}

	if  AdminExist(user.Email) {
		c.AsciiJSON(400,gin.H{
			"msg":"无权删除",
		})
	}else {
		res = DB.Where("id = ?",tipid).Take(&tip)
		if res.Error != nil{fmt.Println(res.Error);return}
		DB.Delete(&tip)
		c.AsciiJSON(400,gin.H{
			"msg":"删除成功",
		})
	}
}

func ViewTip (c *gin.Context){
	DB := database.Link()
	var tips []model.Tip
	res :=DB.Find(&tips)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(tips);i++ {
		var tipimg []model.TipSrc
		var comment []model.TipComment
		res =DB.Where("card_id=?",tips[i].ID).Find(&tipimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		res =DB.Where("card_id=?",tips[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		tips[i].TipSrc = tipimg
		tips[i].TipComment = comment
		//fmt.Println(tips[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": tips,
	})
}


func NewTipComment (c *gin.Context){
	var comment model.TipComment
	err :=c.ShouldBind(&comment)
	if err!=nil{fmt.Println(err);return}
	DB:=database.Link()
	err = DB.Create(&comment).Error
	if err != nil {
		log.Println(err)
	}
	c.AsciiJSON(200,gin.H{
		"msg":"评论成功",
	})
}
func DeleteTipComment (c *gin.Context) {
	var comment model.TipComment
	uid, _ := strconv.Atoi(c.GetHeader("user_id"))//执行删除操作者的id
	commentid := c.GetHeader("comment_id")
	DB := database.Link()
	res := DB.Where("id = ?", commentid).Take(&comment)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	if comment.UserID != uint(uid) {
		c.AsciiJSON(400, gin.H{
			"msg": "无权删除",
		})
	} else {
		DB.Delete(&comment)
		c.AsciiJSON(400, gin.H{
			"msg": "删除成功",
		})
	}
}

func TipCommentLike (c *gin.Context){
	var comment model.TipComment
	commentid :=c.Query("comment_id")
	like :=c.Query("like")
	DB := database.Link()
	res := DB.Where("id = ?",commentid).Take(&comment)
	if res.Error!=nil{fmt.Println(res.Error);return}
	if like == "true"{
		comment.CommentStar+=1
	}else{
		comment.CommentStar-=1
	}
	DB.Save(&comment)
	c.AsciiJSON(200,gin.H{
		"shares": "ok",
	})
}