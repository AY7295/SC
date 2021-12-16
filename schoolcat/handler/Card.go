package handler

import (
	"SchoolCat/database"
	"SchoolCat/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func Card (c *gin.Context){
	var card model.CatCard
	err:=c.ShouldBind(&card)
	if err != nil {
		fmt.Println(err); return
	}
	DB := database.Link()
	//fmt.Println(card)
	err = DB.Create(&card).Error
	if err != nil {
		log.Println(err)
	}
	c.AsciiJSON(200,gin.H{
		"msg":"操作成功",
		"CardID":card.ID,
	})
}

func DeleteCard (c *gin.Context){
	DB := database.Link()
	var card model.CatCard
	cardid := c.Query("card_id")

	var user model.User
	uid,_ := strconv.Atoi(c.Query("user_id"))//执行删除操作者的id
	res := DB.Where("id = ?",uid).Take(&user)
	if res.Error != nil{fmt.Println(res.Error);return}

	if  AdminExist(user.Email) {
		c.AsciiJSON(400,gin.H{
			"msg":"无权删除",
		})
	}else {
		res = DB.Where("id = ?",cardid).Take(&card)
		if res.Error != nil{fmt.Println(res.Error);return}
		DB.Delete(&card)
		c.AsciiJSON(400,gin.H{
			"msg":"删除成功",
		})
	}
}

func ViewCard (c *gin.Context){
	DB := database.Link()
	var cards []model.CatCard
	res :=DB.Find(&cards).Limit(10).Offset(10)
	if res.Error!=nil{fmt.Println(res.Error);return}
	for i:=0;i < len(cards);i++ {
		var cardimg []model.CatCardSrc
		var comment []model.CatCardComment
		res =DB.Where("card_id=?",cards[i].ID).Find(&cardimg)
		if res.Error!=nil{fmt.Println(res.Error);return}
		res =DB.Where("card_id=?",cards[i].ID).Find(&comment)
		if res.Error!=nil{fmt.Println(res.Error);return}
		cards[i].CatCardSrc = cardimg
		cards[i].CatCardComment = comment
		//fmt.Println(shares[i])
	}
	c.AsciiJSON(200,gin.H{
		"shares": cards,
	})
}

func NewCardComment (c *gin.Context){
	var comment model.CatCardComment
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
func DeleteCardComment (c *gin.Context){
	var comment model.CatCardComment
	uid,_ := strconv.Atoi(c.Query("user_id"))
	commentid := c.Query("comment_id")
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