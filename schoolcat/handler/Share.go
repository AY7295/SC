package handler

import (
	"SchoolCat/database"
	"SchoolCat/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func NewShare(c *gin.Context){
	var share model.Share
	err:=c.ShouldBind(&share)
	if err != nil {
		fmt.Println(err); return
	}
	DB := database.Link()
	//fmt.Println(share)
	err = DB.Create(&share).Error
	if err != nil {
		log.Println(err)
	}
	c.AsciiJSON(200,gin.H{
		"msg":"分享成功",
		"ShareID":share.ID,
	})
}
func DeleteShare (c *gin.Context){
	var share model.Share
	uid,err := strconv.Atoi(c.GetHeader("user_id"))
	if err != nil {
		c.AsciiJSON(400,gin.H{
			"msg":"user_id错误",
		})
	}
	shareid := c.GetHeader("share_id")
	DB := database.Link()
	res := DB.Where("id = ?",shareid).Take(&share)
	if res.Error != nil{fmt.Println(res.Error);return}
	if share.UserID !=uint(uid) {
		c.AsciiJSON(400,gin.H{
			"msg":"无权删除",
		})
	}else {
		DB.Delete(&share)
		c.AsciiJSON(400,gin.H{
			"msg":"删除成功",
		})
	}

}
func NewShareComment (c *gin.Context){
	var comment model.UserComment
	err :=c.ShouldBind(&comment)
	if err!=nil{fmt.Println(err);return}
	DB:=database.Link()
	err = DB.Create(&comment).Error
	if err != nil {
		log.Println(err)
	}
	c.AsciiJSON(200,gin.H{
		"msg":"评论成功",
		"user_comment":comment.ID,
	})
}
func DeleteShareComment (c *gin.Context){
	var comment model.UserComment
		uid,_ := strconv.Atoi(c.GetHeader("user_id"))
		commentid := c.GetHeader("comment_id")
		DB := database.Link()
		res := DB.Where("id = ?",commentid).Take(&comment)
		if res.Error != nil{fmt.Println(res.Error);return}
		if comment.UserID !=uint(uid) {
			c.AsciiJSON(400,gin.H{
				"msg":"无权删除",
			})
		}else {
			DB.Delete(&comment)
			c.AsciiJSON(400,gin.H{
				"msg":"删除成功",
			})
		}

}

func Search (c *gin.Context){
	DB := database.Link()
	var shares []model.Share
	keywords :=c.Query("keywords")
	uid := c.GetHeader("user_id")
	res :=DB.Where("content LIKE ?","%"+keywords+"%").Find(&shares).Limit(10).Offset(10)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(shares);i++ {
		var shareimg []model.ShareImage
		var comment []model.UserComment
		var shareLike model.ShareLike

		res =DB.Where("share_id=?",shares[i].ID).Find(&shareimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].ShareImages = shareimg

		res =DB.Where("share_id=?",shares[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		for k:=0;k<len(comment);k++{
			var commentLike model.ShareCommentLike
			res =DB.Where("user_comment_id = ? AND user_id = ?", comment[i].ID, uid).Take(&commentLike)
			if res.Error!=nil{fmt.Println(res.Error);return}
			comment[i].Like=commentLike.Like
		}
		shares[i].UserComment = comment

		res =DB.Where("share_id = ? AND user_id = ?", shares[i].ID, uid).Take(&shareLike)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].Like=shareLike.Like
		//fmt.Println(shares[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": shares,
	})
}

func ViewShare (c *gin.Context){
	DB := database.Link()
	var shares []model.Share
	uid := c.GetHeader("user_id")
	res :=DB.Find(&shares).Limit(10).Offset(10)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(shares);i++ {
		var shareimg []model.ShareImage
		var comment []model.UserComment
		var shareLike model.ShareLike

		res =DB.Where("share_id=?",shares[i].ID).Find(&shareimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].ShareImages = shareimg

		res =DB.Where("share_id=?",shares[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		for k:=0;k<len(comment);k++{
			var commentLike model.ShareCommentLike
			res =DB.Where("user_comment_id = ? AND user_id = ?", comment[i].ID, uid).Take(&commentLike)
			if res.Error!=nil{fmt.Println(res.Error);return}
			comment[i].Like=commentLike.Like
		}
		shares[i].UserComment = comment

		res =DB.Where("share_id = ? AND user_id = ?", shares[i].ID, uid).Take(&shareLike)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].Like=shareLike.Like
		//fmt.Println(shares[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": shares,
	})
}

func SelfShare (c *gin.Context){
	DB := database.Link()
	var shares []model.Share
	uid := c.GetHeader("user_id")
	//fmt.Println(uid)
	res :=DB.Where("user_id = ?",uid).Find(&shares).Limit(10).Offset(10)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(shares);i++ {
		var shareimg []model.ShareImage
		var comment []model.UserComment
		var shareLike model.ShareLike

		res =DB.Where("share_id=?",shares[i].ID).Find(&shareimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].ShareImages = shareimg

		res =DB.Where("share_id=?",shares[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		for k:=0;k<len(comment);k++{
			var commentLike model.ShareCommentLike
			res =DB.Where("user_comment_id = ? AND user_id = ?", comment[i].ID, uid).Take(&commentLike)
			if res.Error!=nil{fmt.Println(res.Error);return}
			comment[i].Like=commentLike.Like
		}
		shares[i].UserComment = comment

		res =DB.Where("share_id = ? AND user_id = ?", shares[i].ID, uid).Take(&shareLike)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].Like=shareLike.Like
		//fmt.Println(shares[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": shares,
	})
}

func ShareCommentLike (c *gin.Context){
	var commentLike model.ShareCommentLike
	var comment model.UserComment
	err :=c.ShouldBind(&commentLike)
	if err!=nil{log.Println(err);return}
	DB := database.Link()
	DB.Create(&commentLike)
	res := DB.Where("id = ?",commentLike.UserCommentID).Take(&comment)
	if res.Error!=nil{fmt.Println(res.Error);return}
	if commentLike.Like == "true"{
		comment.CommentStar+=1
	}else{
		comment.CommentStar-=1
	}
	DB.Save(&comment)
	c.AsciiJSON(200,gin.H{
		"shares": "ok",
	})
}

func ShareLike (c *gin.Context){

	var shareLike,shareLike0 model.ShareLike
	var share model.Share
	err :=c.ShouldBind(&shareLike)
	if err!=nil{log.Println(err);return}
	DB := database.Link()
	res := DB.Where("id = ?",shareLike.ShareID).Take(&share)
	if res.Error!=nil{fmt.Println(res.Error);return}
	if shareLike.Like == "true" {
		share.ShareStar+=1
	}else{
		share.ShareStar-=1
	}
	DB.Save(&share)
	res = DB.Where("id = ?",shareLike.ID).Take(&shareLike0)
	if res.RowsAffected==0{
		DB.Create(&shareLike)
	}else {
		shareLike0.Like = shareLike.Like
		DB.Save(&shareLike0)
	}
	c.AsciiJSON(200,gin.H{
		"shares": "ok",
	})

}