package midware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func CORS() gin.HandlerFunc {//跨域用的暂时没写好

	//return func(c *gin.Context) {
	//	c.Header("Access-Control-Allow-Origin", "*")
	//	c.Header("Access-Control-Allow-Credentials", "true")
	//	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-CSRF-Token, X-Requested-With, User-Agent, DNT, Keep-Alive")
	//	c.Header("Access-Control-Allow-Methods", "GET, PUT, PATCH, DELETE, OPTIONS, POST")
	//	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Access-Control-Allow-Headers")
	//	if method := c.Request.Method; method == "OPTIONS" {
	//	c.AbortWithStatus(http.StatusNoContent)
	//	}
	//}

	return func(c *gin.Context){
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, PUT, PATCH, DELETE, OPTIONS, POST")
		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type,id,comment_id,user_id,auth,card_id")
		if method := c.Request.Method; method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}

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
