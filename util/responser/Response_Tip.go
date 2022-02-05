package responser

import (
	"SchoolCat/model"
	"github.com/gin-gonic/gin"
)

func TipSucceed(c *gin.Context, tipId uint) {
	c.AsciiJSON(200, gin.H{
		"msg":   "操作成功",
		"TipID": tipId,
	})
}

func DisplayTips(c *gin.Context, tips []model.Tip) {
	c.AsciiJSON(200, gin.H{
		"tips": tips,
	})
}
