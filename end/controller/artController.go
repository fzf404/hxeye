package controller

import (
	"App/database"
	"App/dto"
	"App/model"
	"App/response"

	"github.com/gin-gonic/gin"
)

// AddArticle 添加文章
func AddArticle(ctx *gin.Context) {
	db := database.GetDB()

	// 获得文章信息及用户信息
	user, _ := ctx.Get("user")
	article, _ := ctx.Get("article")

	newArt := model.Article{
		UserID:   user.(model.User).ID,
		UserName: user.(model.User).Name,
		Title:    article.(model.Article).Title,
		ArtType:  article.(model.Article).ArtType,
		Content:  article.(model.Article).Content,
	}

	db.Create(&newArt)
	response.Success(ctx, gin.H{"name": user.(model.User).Name, "title": newArt.Title, "artid": newArt.ID}, "文章保存成功")
}

// DelArticle 删除文章
func DelArticle(ctx *gin.Context) {
	db := database.GetDB()

	user, _ := ctx.Get("user")

	// 获取用户即文章信息
	artID := ctx.PostForm("artid")

	if len(artID) == 0 {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	var article model.Article
	db.First(&article, artID)

	if article.ID == 0 {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	db.Delete(&article)
	response.Success(ctx, gin.H{"name": user.(model.User).Name, "title": article.Title, "artid": article.ID}, "删除成功")
}

// ModArticle 修改文章
func ModArticle(ctx *gin.Context) {
	db := database.GetDB()

	user, _ := ctx.Get("user")
	artID := ctx.PostForm("artid")

	if len(artID) == 0 {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 查找文章
	var art model.Article
	db.First(&art, artID)

	if art.ID == 0 {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	article, _ := ctx.Get("article")

	art.Title = article.(model.Article).Title
	art.ArtType = article.(model.Article).ArtType
	art.Content = article.(model.Article).Content

	db.Save(&art)
	response.Success(ctx, gin.H{"name": user.(model.User).Name, "title": art.Title, "artid": art.ID}, "文章更新成功")
}

// PublicArticle 发布文章
func PublicArticle(ctx *gin.Context) {
	db := database.GetDB()

	artID := ctx.PostForm("artid")

	if len(artID) == 0 {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 查找文章
	var art model.Article
	db.First(&art, artID)

	if art.ID == 0 {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	db.Model(&art).Update("public", 1)

	response.Success(ctx, gin.H{"title": art.Title, "artid": art.ID}, "文章发布成功")
}

// GetSecretArticle 获取文章
func GetSecretArticle(ctx *gin.Context) {
	db := database.GetDB()

	// 获取文章
	artID := ctx.PostForm("artid")
	var art model.Article
	db.First(&art, artID)

	if art.ID == 0 {
		response.NotFind(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"article": dto.ArticleDto(art)}, "获取成功")
}
