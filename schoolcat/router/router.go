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

	user.GET("/login",handler.Login)//登录
	user.POST("/register",handler.Register)//注册
	user.POST("/info",handler.AddInfo)//添加信息：邮箱，性别，学校，昵称
	user.PUT("/info",handler.UpdateInfo)//传过来信息和添加一样


	//user.GET("/share",handler.MoreShare)//刷新请求新的share


	user.POST("newShare",handler.NewShare)//用户添加share
	user.POST("/")

	user.GET("/viewShare",handler.ViewShare)
	//user.GET("/viewShareSrc",handler.ViewShareSrc)

	//user.GET("/search",handler.Search)//用户搜索


	return engine
}
