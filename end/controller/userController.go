package controller

import (
	"App/database"
	"App/dto"
	"App/middleware"
	"App/model"
	"App/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login 用户登录信息处理
func Login(ctx *gin.Context) {
	db := database.GetDB()

	//获取数据
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	// 判断密码
	if len(password) < 8 || len(password) > 20 {
		response.Fail(ctx, nil, "用户名或密码错误")
		return
	}

	// 判断用户是否存在
	var user model.User
	db.Where("name = ?", name).First(&user)
	// 用户名判断
	if user.ID == 0 {
		response.Fail(ctx, nil, "用户名或密码错误")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "用户名或密码错误")
		return
	}
	// 分发Token
	token, err := middleware.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "Token发放失败")
		log.Print("token generate error:", err)
		return
	}

	response.Success(ctx, gin.H{"token": token, "user": name}, "登陆成功")
}

// 修改密码
func ChangePwd(ctx *gin.Context) {

	db := database.GetDB()

	// 获得用户数据
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	if len(password) < 8 || len(password) > 20 {
		response.Warning(ctx, nil, "密码须在8-20位之间")
		return
	}

	user, _ := ctx.Get("user")
	if user.(model.User).Name != name && !user.(model.User).Super {
		response.Fail(ctx, nil, "修改请求非法")
		return
	}

	// 创建新密码
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}
	// 更新密码并保存
	modUser := user.(model.User)
	modUser.Password = string(hashPassword)
	db.Save(modUser)

	response.Success(ctx, nil, "密码修改成功")
}

// AddUser 新建用户
func AddUser(ctx *gin.Context) {
	db := database.GetDB()

	// 获得新增用户数据
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	if len(password) < 8 || len(password) > 20 {
		response.Warning(ctx, nil, "密码须在8-20位之间")
		return
	}

	// 判断用户名是否存在
	var user model.User
	db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		response.Warning(ctx, nil, "用户名已存在")
		return
	}

	// 创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}
	newUser := model.User{
		Name:     name,
		Password: string(hashPassword),
		Super:    false,
	}
	db.Create(&newUser)

	response.Success(ctx, nil, "注册成功")
}

// MyInfo 输出我的信息
func UserInfo(ctx *gin.Context) {
	// 获取中间件添加的user
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.TouserMyDto(user.(model.User))}, "个人信息获取成功")
}
