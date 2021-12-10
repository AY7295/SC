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
	uid,_ := strconv.Atoi(c.GetHeader("user_id"))
	commentid := c.GetHeader("share_id")
	DB := database.Link()
	res := DB.Where("id = ?",commentid).Take(&share)
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
	res :=DB.Where("content LIKE ?","%"+keywords+"%").Find(&shares).Limit(10).Offset(10)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(shares);i++ {
		var shareimg []model.ShareImage
		var comment []model.UserComment
		res =DB.Where("share_id=?",shares[i].ID).Find(&shareimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		res =DB.Where("share_id=?",shares[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].ShareImages = shareimg
		shares[i].UserComment = comment
		//fmt.Println(shares[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": shares,
	})
}

func ViewShare (c *gin.Context){
	DB := database.Link()
	var shares []model.Share
	res :=DB.Find(&shares).Limit(10).Offset(10)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(shares);i++ {
		var shareimg []model.ShareImage
		var comment []model.UserComment
		res =DB.Where("share_id=?",shares[i].ID).Find(&shareimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		res =DB.Where("share_id=?",shares[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].ShareImages = shareimg
		shares[i].UserComment = comment
		//fmt.Println(shares[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": shares,
	})
}

func SelfShare (c *gin.Context){
	DB := database.Link()
	var shares []model.Share
	id := c.GetHeader("UID")
	res :=DB.Where("user_id = ?",id).Find(&shares).Limit(10).Offset(10)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(shares);i++ {
		var shareimg []model.ShareImage
		var comment []model.UserComment
		res =DB.Where("share_id=?",shares[i].ID).Find(&shareimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		res =DB.Where("share_id=?",shares[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		shares[i].ShareImages = shareimg
		shares[i].UserComment = comment
		//fmt.Println(shares[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": shares,
	})
}

func CommentLike (c *gin.Context){
	var comment model.UserComment
	commentid :=c.Query("comment_id")
	like :=c.Query("like")
	DB := database.Link()
	res := DB.Where("id = ?",commentid).Take(&comment)
	if res.Error!=nil{fmt.Println(res.Error);return}
	if like == "yes"{
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
	var share model.Share
	shareid :=c.Query("share_id")
	like :=c.Query("like")
	DB := database.Link()
	res := DB.Where("id = ?",shareid).Take(&share)
	if res.Error!=nil{fmt.Println(res.Error);return}
	if like == "yes"{
		share.ShareStar+=1
	}else{
		share.ShareStar-=1
	}
	DB.Save(&share)
	c.AsciiJSON(200,gin.H{
		"shares": "ok",
	})
}