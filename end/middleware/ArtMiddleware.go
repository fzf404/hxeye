package middleware

import (
	"App/model"
	"App/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArtMiddleware 判断文章是否合法
func ArtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.PostForm("title")

		if len(title) < 1 {
			response.Warning(ctx, nil, "标题不能为空")
			ctx.Abort()
			return
		}

		artType, err := strconv.Atoi(ctx.PostForm("type"))

		if err != nil || artType < 1 || artType > 3 {
			response.Warning(ctx, nil, "文章类型必须为1-3")
			ctx.Abort()
			return
		}

		content := ctx.PostForm("content")

		if len(content) < 1 {
			response.Warning(ctx, nil, "文章不能为空")
			ctx.Abort()
			return
		}

		article := model.Article{
			Title:   title,
			ArtType: artType,
			Content: content,
		}
		ctx.Set("article", article)
		ctx.Next()
	}
}
