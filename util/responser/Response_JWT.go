package responser

import "github.com/gin-gonic/gin"

func InvalidToken(c *gin.Context) {
	c.AsciiJSON(400, gin.H{
		"msg": "无效的Token",
	})
}

func OverTimedToken(c *gin.Context) {
	c.AsciiJSON(400, gin.H{
		"msg": "token失效，请重新登录",
	})
}

func WrongToken(c *gin.Context) {
	c.AsciiJSON(400, gin.H{
		"msg": "Token错误",
	})
}

func IllegalAccess(c *gin.Context) {
	c.AsciiJSON(400, gin.H{
		"msg": "非法访问",
	})
}
