package handler

import (
	"SchoolCat/database"
	"SchoolCat/model"
	response "SchoolCat/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func NewCard (c *gin.Context){
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
	response.CardSucceed(c,card.ID)
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
		response.UserIdWrong(c)
	}else {
		res = DB.Where("id = ?",cardid).Take(&card)
		if res.Error != nil{fmt.Println(res.Error);return}
		DB.Delete(&card)
		response.DeleteSucceed(c)
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
	response.DisplayCards(c,cards)
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
	response.CommentSucceed(c,comment.ID)
}
func DeleteCardComment (c *gin.Context){
	var comment model.CatCardComment
	uid,_ := strconv.Atoi(c.Query("user_id"))
	commentid := c.Query("comment_id")
	DB := database.Link()
	res := DB.Where("id = ?",commentid).Take(&comment)
	if res.Error != nil{fmt.Println(res.Error);return}
	if comment.UserID !=uint(uid) {
		response.UserIdWrong(c)
	}else {
		DB.Delete(&comment)
		response.DeleteSucceed(c)
	}
}