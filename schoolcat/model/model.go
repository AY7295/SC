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
}

type Admin struct{
	Email   string `form:"email,omitempty" json:"email,omitempty"  gorm:"size:50;unique"` //email可以是任何形式的string，但长度不能超过50
	Password   string `form:"pwd,omitempty" json:"pwd,omitempty" binding:"max=16" gorm:"size:16"` //密码最大16位，前端可以设置一下最小6位
}

type Share struct { //分享
	gorm.Model        //这个id用来标记每个图片的位置
	UserID     uint    `form:"uid" json:"uid" binding:"required" ` //User.ID
	Address    string `form:"address" json:"address" binding:"required" gorm:"default:'null'"` //存地址
	Title      string `form:"title" json:"title" binding:"required" gorm:"not null"`
	Content    string `form:"content" json:"content" binding:"" gorm:"type:longtext"` //附加内容
	//Star uint
}

type Tips struct { //贴士
	gorm.Model        //这个id用来标记每个图片的位置
	Username   string `form:"username" json:"username" binding:"required" gorm:"size:20;not null"`
	Title      string `form:"title" json:"title" binding:"required" gorm:"not null"`
	Content    string `form:"content" json:"content" binding:"required" gorm:"type:longtext"` //附加内容
}

type ShareImg struct { //存share图片地址
	gorm.Model
	ShareID uint    //查图片的时候根据where("ShareID=?")来查
	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}

type TipImg struct { //存tip图片地址
	gorm.Model
	TipID uint    //查图片的时候根据where("TipID=?")来查
	Src   string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}
