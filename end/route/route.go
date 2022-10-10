package route

import (
	"App/controller"
	"App/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由汇总
func CollectRoute(r *gin.Engine) *gin.Engine {

	// 跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 首页信息
	r.GET("/active", controller.GetActive)
	// 登陆
	r.POST("/login", controller.Login)
	// 获取文章
	r.POST("/getart", controller.GetArticle)
	// 获取文章
	r.POST("/artlist", controller.GetArticleList)
	// 搜索
	r.POST("/search", controller.Search)

	adminRouters := r.Group("")
	adminRouters.Use(middleware.AuthMiddleware())

	// 用户信息
	adminRouters.POST("/userinfo", controller.UserInfo)
	// 上传图片
	adminRouters.POST("/upimg", controller.UploadImg)
	// 全部文章
	adminRouters.POST("/allart", controller.GetAllArticle)
	// 获得未发布文章
	adminRouters.POST("/secretart", controller.GetSecretArticle)
	// 添加文章
	adminRouters.POST("/addart", middleware.ArtMiddleware(), controller.AddArticle)
	// 修改文章
	adminRouters.POST("/modart", middleware.ArtMiddleware(), controller.ModArticle)

	// 增加用户
	adminRouters.POST("/adduser", middleware.RootAuthMiddleware(), controller.AddUser)
	// 修改密码
	adminRouters.POST("/modpwd", middleware.RootAuthMiddleware(), controller.ChangePwd)
	// 删除文章
	adminRouters.POST("/delart", middleware.RootAuthMiddleware(), controller.DelArticle)
	// 发布文章
	adminRouters.POST("/pubart", middleware.RootAuthMiddleware(), controller.PublicArticle)
	// 修改首页信息
	adminRouters.POST("/modinfo", middleware.RootAuthMiddleware(), controller.ModInfo)

	return r
}
