package model

import "gorm.io/gorm"

type Share struct { //分享
	gorm.Model        //这个id用来标记每个图片评论的位置:"shareid",这个creatat用来标记maoreshare的时间
	UserID     uint    `json:"UID" ` //User.ID
	Username string `json:"username"`
	Iocnsrc string	`json:"iconsrc" gorm:"unique"`
	Address    string `form:"address" json:"address" gorm:"default:'null'"` //存地址
	//Title      string `form:"title" json:"title" binding:"required" gorm:"size:30"`//标题长度不超过30
	Content    string `form:"content" json:"content"  gorm:"type:longtext"` //附加内容
	ShareStar uint	`json:"ShareStar"`
	ShareImages []ShareImage `json:"share_images"`
	UserComment []UserComment `json:"user_comment"`
}
type ShareImage struct { //存share图片地址
	gorm.Model
	ShareID uint   `json:"share_id"` //查图片的时候根据where("ShareID=?")来查
	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}

type UserComment struct {
	gorm.Model
	Username string	`json:"username" gorm:"size:80"`
	Iocnsrc string	`json:"iconsrc" `
	UserID     uint    `json:"UID" binding:"required" ` //User.ID
	ShareID uint   `json:"share_id"` //查图片的时候根据where("ShareID=?")来查
	CommentStar uint	`json:"commentStar"`
	Comment string `json:"comment"`
	//CommentSrc []CommentSrc `json:"comment_src"`
}

//type CommentSrc struct {
//	gorm.Model
//	CommentID uint   `json:"comment_id"` //查图片的时候根据where("ShareID=?")来查
//	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
//}