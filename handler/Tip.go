package handler

import (
	"SchoolCat/model"
	response "SchoolCat/util"
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

	//fmt.Println(tip)
	err = DB.Create(&tip).Error
	if err != nil {
		log.Println(err)
	}
	response.TipSucceed(c,tip.ID)
}

func DeleteTip(c *gin.Context){

	var tip model.Tip
	tipId := c.Query("tip_id")

	var admin model.Admin
	uid,_ := strconv.Atoi(c.Query("user_id"))//执行删除操作者的id
	res := DB.Where("user_id = ?",uid).Take(&admin)
	if res.Error != nil{fmt.Println(res.Error);return}

	if  res.RowsAffected==0 {
		response.UserIdWrong(c)
	}else {
		res = DB.Where("id = ?", tipId).Take(&tip)
		if res.Error != nil{fmt.Println(res.Error);return}
		DB.Delete(&tip)
		response.DeleteSucceed(c)
	}
}

func ViewTip (c *gin.Context){

	var tips []model.Tip
	uid := c.Query("user_id")
	res :=DB.Find(&tips)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(tips);i++ {
		var tipimg []model.TipSrc
		res =DB.Where("tip_id=?",tips[i].ID).Find(&tipimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		tips[i].TipSrc = tipimg

		var comment []model.TipComment
		res =DB.Where("tip_id=?",tips[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		for k:=0;k<len(comment);k++{
			var commentLike model.ShareCommentLike
			res =DB.Where("user_comment_id = ? AND user_id = ?", comment[i].ID, uid).Take(&commentLike)
			if res.Error!=nil{fmt.Println(res.Error);return}
			comment[i].Like=commentLike.Like
		}
		tips[i].TipComment = comment
		//fmt.Println(tips[i])
	}
	response.DisplayTips(c,tips)
}


func NewTipComment (c *gin.Context){
	var comment model.TipComment
	err :=c.ShouldBind(&comment)
	if err!=nil{fmt.Println(err);return}

	err = DB.Create(&comment).Error
	if err != nil {
		log.Println(err)
	}
	response.CommentSucceed(c,comment.ID)
}
func DeleteTipComment (c *gin.Context) {
	var comment model.TipComment
	uid, _ := strconv.Atoi(c.Query("user_id"))//执行删除操作者的id
	commentid := c.Query("comment_id")

	res := DB.Where("id = ?", commentid).Take(&comment)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	if comment.UserID != uint(uid) {
		response.UserIdWrong(c)
	} else {
		DB.Delete(&comment)
		response.DeleteSucceed(c)
	}
}

func TipCommentLike (c *gin.Context){
	var comment model.TipComment
	commentid :=c.Query("comment_id")
	like :=c.Query("like")

	res := DB.Where("id = ?",commentid).Take(&comment)
	if res.Error!=nil{fmt.Println(res.Error);return}
	if like == "true"{
		comment.CommentStar+=1
	}else{
		comment.CommentStar-=1
	}
	DB.Save(&comment)
	response.Like(c)
}
