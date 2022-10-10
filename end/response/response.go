package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 封装响应信息
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// Success 成功响应
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// Warning 提示响应
func Warning(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 422, data, msg)
}

// NotFind 找不到响应
func NotFind(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 404, data, msg)
}

// Fail 错误响应
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
