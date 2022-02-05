package router

import (
	"SchoolCat/config"
	"SchoolCat/handler/service"
	"SchoolCat/midware"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Router() {
	gin.SetMode(config.AppMode)

	engine := gin.Default()
	engine.Use(midware.CORS())

	engine.POST("/login", service.Login)       //登录
	engine.POST("/register", service.Register) //注册

	admin := engine.Group("/admin", midware.Admin(), midware.JWT())
	user := engine.Group("/user", midware.JWT())

	//百科相关
	admin.POST("/tip", service.Tip)         //add tip
	admin.DELETE("/tip", service.DeleteTip) //delete tip

	user.POST("/newTipComment", service.NewTipComment) //用户添加评论
	user.DELETE("/deleteTipComment", service.DeleteTipComment)
	user.GET("/newTip", service.ViewTip)
	user.PUT("/tipCommentLike", service.TipCommentLike) //用户点赞

	//资料卡片相关
	admin.POST("/newCard", service.NewCard)         //add card
	admin.DELETE("/deleteCard", service.DeleteCard) //delete card

	user.POST("/newCArdComment", service.NewCardComment) //用户添加评论
	user.DELETE("/deleteCardComment", service.DeleteCardComment)
	user.GET("newCard", service.ViewCard)

	//用户相关
	user.POST("/info", service.Info) //添加更改信息：邮箱，性别，学校，昵称，简介

	//share相关
	user.GET("/selfShare", service.SelfShare) //用户界面刷新自己的分享
	user.GET("/viewShare", service.ViewShare) //用户请求share

	user.POST("newShare", service.NewShare) //用户添加share
	user.DELETE("/deleteShare", service.DeleteShare)
	user.POST("/newShareComment", service.NewShareComment) //用户添加评论
	user.DELETE("/deleteShareComment", service.DeleteShareComment)

	user.GET("/search", service.Search) //用户搜索

	user.PUT("/shareCommentLike", service.ShareCommentLike) //用户点赞
	user.PUT("/shareLike", service.ShareLike)               //用户点赞

	err := engine.Run(config.HttpPort)
	if err != nil {
		fmt.Println("路由运行端口出错", err)
	}
}
