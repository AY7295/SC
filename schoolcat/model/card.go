package model

import "gorm.io/gorm"

type CatCard struct {
	gorm.Model
	CatName string `json:"cat_name"`
	FurColor string `json:"fur_color"`
	Health string `json:"health"`
	Ster string `json:"ster"`//绝育情况
	Area string `json:"area"`//出没范围·
	Appearance string `json:"appearance"`
	Time string `json:"time"`//第一次目击时间
	Relationship	string `json:"relationship"`//猫际关系
	CatCardComment []CatCardComment
	CatCardSrc []CatCardSrc	`json:"cat_card_src"`
}
type CatCardSrc struct {
	gorm.Model
	CatCardID uint   `json:"cat_card_id"` //查图片的时候根据where("ShareID=?")来查
	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址
}
type CatCardComment struct {
	gorm.Model
	CommentID uint   `json:"comment_id"` //查图片的时候根据where("ShareID=?")来查
	Src     string `form:"src" json:"src" binding:"required" gorm:"type:longtext"` //图片地址

}