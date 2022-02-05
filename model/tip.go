package model

import "gorm.io/gorm"

type Tip struct {
	gorm.Model
	CreateId   uint         `json:"create_id"`
	DeleteID   uint         `json:"delete_id"`
	Username   string       `json:"username" gorm:"size:80"` //创建者
	IconSrc    string       `json:"icon_src" `
	ModelCode  int          `json:"model_code"` //根据modelCode来区别tip类型
	Title      string       `json:"title"`
	Content    string       `json:"content" gorm:"type:longtext"`
	TipComment []TipComment `json:"tip_comment"`
	TipSrc     []TipSrc     `json:"tip_src"`
}
type TipSrc struct {
	gorm.Model
	TipID uint   `json:"card_id"`
	Src   string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}
type TipComment struct {
	gorm.Model
	Username    string `json:"username" gorm:"size:80"`
	IconSrc     string `json:"icon_src" `
	UserID      uint   `json:"user_id" binding:"required" ` //User.ID
	TipID       uint   `json:"tip_id"`
	Like        string `json:"like"`
	CommentStar uint   `json:"comment_star"`
	Comment     string `json:"comment"`
}

type TipCommentLike struct {
	UserID       uint   `json:"user_id" binding:"required" ` //User.ID
	TipCommentID uint   `json:"tip_comment_id"`
	Like         string `json:"like"`
}
