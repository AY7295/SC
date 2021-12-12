package midware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func CORS() gin.HandlerFunc {//跨域用的暂时没写好
	//
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

/*const MySecret = "1qaz2wsx"
type MyClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("email", mc.Email)
		c.Next() // 后续的处理函数可以用过c.Get("email")来获取当前请求的用户信息
	}
}

// GenToken 生成JWT
func GenToken(email string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		email, // 自定义字段
		//Passwaord
		jwt.StandardClaims{
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
*/


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
