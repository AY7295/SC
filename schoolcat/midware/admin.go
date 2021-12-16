package midware

import "github.com/gin-gonic/gin"

func Admin() gin.HandlerFunc {//认证管理员

	return func(c *gin.Context) {
		auth := c.GetHeader("auth") //通过Header中的key对应的值来判断管理员与
		if auth != "12345"{
			c.AbortWithStatusJSON(401, gin.H{
				"message": "无权访问",
			})
			return
		}
		c.Next()
	}
}
