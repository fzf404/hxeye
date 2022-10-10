package middleware

import (
	"App/database"
	"App/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 验证token以及信息返回
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取 Token
		tokenString := ctx.GetHeader("Authorization")

		// 判断Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 判断Token
		tokenString = tokenString[7:]
		token, claims, err := ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// token通过验证, 获取claims中的UserID
		userID := claims.UserID
		db := database.GetDB()
		var user model.User
		db.First(&user, userID)

		// 验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 用户存在 将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}

// AuthMiddleware 验证token以及信息返回
func RootAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		isSuper := user.(model.User).Super
		if !isSuper {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
