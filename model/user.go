package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

// Claim JWT申明字段
type Claim struct {
	Email string
	jwt.StandardClaims
}

type User struct {
	gorm.Model        // 这个id用来标记每个人的share
	Email      string `json:"email,omitempty"  gorm:"size:50;unique"`      //email可以是任何形式的string，但长度不能超过50
	Password   string `json:"pwd,omitempty" binding:"max=16" `             //密码最大16位，前端可以设置一下最小6位
	Username   string `json:"username,omitempty" gorm:"size:20"`           //username可以是任何形式的string，但长度不能超过20,必须要有，用户没填前端就随机给一个
	Gender     string `json:"gender,omitempty" gorm:"default:'保密';size:2"` //性别只传男或女，空值默认为保密
	School     string `son:"school,omitempty" gorm:"default:'null';size:50"`
	Resume     string `json:"resume,omitempty"  gorm:"default:'null',type:longtext"`
	IconSrc    string `json:"icon_src,omitempty" gorm:"default:'null'"` //头像
}

type Admin struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Email    string `form:"email,omitempty" json:"email,omitempty"  gorm:"size:50;unique"`      //email可以是任何形式的string，但长度不能超过50
	Password string `form:"pwd,omitempty" json:"pwd,omitempty" binding:"max=16" gorm:"size:16"` //密码最大16位，前端可以设置一下最小6位
}
