package model

import "gorm.io/gorm"

type Tip struct {
	gorm.Model
	CreateId   uint         `json:"create_id"`
	DeleteID   uint         `json:"delete_id"`
	Username string	`json:"username" gorm:"size:80"`//创建者
	Iocnsrc string	`json:"iconsrc" `
	ModelCode int `json:"model_code"`//根据modelcode来区别tip类型
	Title      string       `json:"title"`
	Content    string       `json:"content" gorm:"type:longtext"`
	TipComment []TipComment `json:"cat_card_comment"`
	TipSrc     []TipSrc     `json:"cat_card_src"`
}
type TipSrc struct {
	gorm.Model
	TipID uint   `json:"card_id"`
	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}
type TipComment struct {
	gorm.Model
	Username string	`json:"username" gorm:"size:80"`
	Iocnsrc string	`json:"iconsrc" `
	UserID     uint    `json:"UID" binding:"required" ` //User.ID
	TipID uint   `json:"card_id"`
	CommentStar uint	`json:"commentStar"`
	Comment string `json:"comment"`
}