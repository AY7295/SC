package handler

import (
	"SchoolCat/database"
	"SchoolCat/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func NewShare(c *gin.Context){
	var share model.Share
	err:=c.ShouldBind(&share)
	if err != nil {
		fmt.Println(err); return
	}
	DB := database.Link()
	fmt.Println(share)
	err = DB.Create(&share).Error
	if err != nil {
		log.Println(err)
	}
	c.AsciiJSON(200,gin.H{
		"msg":"分享成功",
		"ShareID":share.ID,
	})
}
func NewComment (c *gin.Context){
	var comment model.UserComment
	err :=c.ShouldBind(&comment)
	if err!=nil{fmt.Println(err);return}
}

func ViewShare(c *gin.Context){
	id := c.Query("shareid")
	var share model.Share
	var shareimg []model.ShareImage
	DB := database.Link()
	res := DB.Where("id = ?", id).Take(&share)
	if res.Error!=nil{fmt.Println(res.Error);return}
	res =DB.Where("share_id=?",id).Find(&shareimg)
	if res.Error!=nil{fmt.Println(res.Error);return}
	share.ShareImages=shareimg
	c.AsciiJSON(200,gin.H{
		"share":share,
	})
}

