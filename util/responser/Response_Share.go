package responser

import (
	"SchoolCat/model"
	"github.com/gin-gonic/gin"
)

func ShareSucceed(c *gin.Context, shareId uint) {
	c.AsciiJSON(200, gin.H{
		"msg":     "操作成功",
		"ShareID": shareId,
	})
}

func UserIdWrong(c *gin.Context) {
	c.AsciiJSON(400, gin.H{
		"msg": "user_id错误",
	})
}

func DeleteSucceed(c *gin.Context) {
	c.AsciiJSON(400, gin.H{
		"msg": "删除成功",
	})
}

func CommentSucceed(c *gin.Context, commentId uint) {
	c.AsciiJSON(200, gin.H{
		"msg":       "评论成功",
		"commentID": commentId,
	})
}

func DisplayShares(c *gin.Context, shares []model.Share) {
	c.AsciiJSON(200, gin.H{
		"shares": shares,
	})
}

func Like(c *gin.Context) {
	c.AsciiJSON(200, gin.H{
		"msg": "ok",
	})
}
