package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // 这个id用来标记每个人的share
	Email   string `form:"email,omitempty" json:"email,omitempty"  gorm:"size:50;unique"` //email可以是任何形式的string，但长度不能超过50
	Password   string `form:"pwd,omitempty" json:"pwd,omitempty" binding:"max=16" gorm:"size:16"` //密码最大16位，前端可以设置一下最小6位
	Username	string `form:"username,omitempty" json:"username,omitempty" gorm:"size:20"`//username可以是任何形式的string，但长度不能超过20
	Gender     string `form:"gender,omitempty" json:"gender,omitempty" gorm:"default:'保密';size:2"` //性别只传男或女，空值默认为保密
	School     string `form:"school,omitempty" json:"school,omitempty" gorm:"default:'null';size:50"`
	Resume string`json:"resume"  gorm:"default:'null'"`
	IconSrc string `json:"icon_src" gorm:"default:'null'"`//头像

}

type Admin struct{
	Email   string `form:"email,omitempty" json:"email,omitempty"  gorm:"size:50;unique"` //email可以是任何形式的string，但长度不能超过50
	Password   string `form:"pwd,omitempty" json:"pwd,omitempty" binding:"max=16" gorm:"size:16"` //密码最大16位，前端可以设置一下最小6位
}
