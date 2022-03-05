package midware

import (
	"SchoolCat/config"
	"SchoolCat/database"
	"SchoolCat/model"
	response "SchoolCat/util/responser"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	jwtKey       = []byte(config.C.Token.JwtKey)
	expireMinute = config.C.Token.ExpireMinute
)

func GenerateToken(email string) string {

	var claims = model.Claim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                               //立即生效
			ExpiresAt: time.Now().Add(expireMinute * time.Hour).Unix(), //失效时间
			Issuer:    "test-sever",                                    //签发者
		},
	}

	res := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	Token, err := res.SignedString(jwtKey)
	if err != nil {
		fmt.Println("err:", err)
		return ":"
	}

	//fmt.Println(Token)
	return "MaoMao " + Token
}

func ParseToken(tokenString string) (claim *model.Claim) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return string(jwtKey), nil
	})

	if err != nil {
		fmt.Println(err)
	}
	return token.Claims.(*model.Claim)
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("token")
		uid := c.Query("user_id")
		//fmt.Println(authHeader,uid)
		if authHeader == "" {
			response.IllegalAccess(c)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		//parts := strings.Fields(authHeader)
		fmt.Println(parts)
		if parts[0] != "MaoMao" || len(parts) != 2 {
			response.WrongToken(c)
			c.Abort()
			return
		}
		claim := ParseToken(parts[1])

		var user model.User
		database.DB.Where("id = ?", uid).Take(&user)
		if user.Email != claim.Email {
			response.InvalidToken(c)
			c.Abort()
			return
		}
		if time.Now().Unix() > claim.ExpiresAt {
			response.OverTimedToken(c)
			c.Abort()
			return
		}
		//更新token
		//c.Set("token",GenerateToken(user.Email))
		//c.AsciiJSON(200,gin.H{
		//	"token":GenerateToken(claim.Email),
		//})
		c.Next()

	}
}
