package router

import (
	"SchoolCat/handler"
	"SchoolCat/midware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine  {
	engine := gin.Default()
	engine.Use(midware.CORS())
	admin := engine.Group("/admin",midware.Admin())
	user :=	engine.Group("/user")

	admin.POST("/tip",handler.AddTip)//add tip
	admin.PUT("/tip",handler.UpdateTip)//update tip
	admin.DELETE("/tip",handler.DeleteTip)//delete tip

	user.POST("/login",handler.Login)//登录
	user.POST("/register",handler.Register)//注册
	user.POST("/info",handler.Info)//添加更改信息：邮箱，性别，学校，昵称，简介

	user.POST("newShare",handler.NewShare)//用户添加share
	user.POST("/newComment",handler.NewComment)//用户添加评论
	user.GET("/viewShare",handler.ViewShare)//用户请求share


	//user.GET("/search",handler.Search)//用户搜索


	return engine
}
