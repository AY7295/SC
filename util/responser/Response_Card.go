package responser

import (
	"SchoolCat/model"
	"github.com/gin-gonic/gin"
)

func DisplayCards(c *gin.Context, cards []model.CatCard) {
	c.AsciiJSON(200, gin.H{
		"cards": cards,
	})
}

func CardSucceed(c *gin.Context, cardId uint) {
	c.AsciiJSON(200, gin.H{
		"msg":    "操作成功",
		"CardID": cardId,
	})
}
