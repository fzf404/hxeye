package controller

import (
	"App/database"
	"App/dto"
	"App/model"
	"App/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// GetArticle 获取文章
func GetArticle(ctx *gin.Context) {
	db := database.GetDB()

	// 获取文章
	artID := ctx.PostForm("artid")
	var art model.Article
	db.First(&art, artID)

	if art.ID == 0 || !art.Public {
		response.NotFind(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"article": dto.ArticleDto(art)}, "获取成功")
}

// GetArticleList 全部可查看文章列表
func GetArticleList(ctx *gin.Context) {

	db := database.GetDB()

	// 文章信息处理
	pageid := ctx.PostForm("pageid")
	if len(pageid) == 0 {
		pageid = "1"
	}

	// 获得pageid
	cpageid, err := strconv.Atoi(pageid)

	if err != nil {
		response.Fail(ctx, nil, "请求字段非法")
		return
	}

	// map处理全部articles
	var articles []model.Article
	items := make(map[string]model.ArticleDto)

	getmaxart := viper.GetString("common.listart")
	maxart, _ := strconv.Atoi(getmaxart)
	// 获取列表
	db.Where("public = ?", 1).Order("id desc").Limit(maxart).Offset(maxart * (cpageid - 1)).Find(&articles)
	// 判断获取情况
	if len(articles) == 0 {
		response.NotFind(ctx, nil, "未找到已发布文章")
		return
	}

	for index, art := range articles {
		// 使用dto处理全部article
		items["art"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}

	response.Success(ctx, gin.H{
		"pageid": pageid,
		"arts":   items,
	}, "获取全部文章成功")
}

// GetAllArticle 全部文章列表
func GetAllArticle(ctx *gin.Context) {
	db := database.GetDB()

	// 文章信息处理
	pageid := ctx.PostForm("pageid")
	hide := ctx.PostForm("hide")
	if len(pageid) == 0 {
		pageid = "1"
	}

	cpageid, err := strconv.Atoi(pageid)

	if err != nil {
		response.Fail(ctx, nil, "请求字段非法")
		return
	}

	// map处理全部articles
	var articles []model.Article
	items := make(map[string]model.ArticleDto)

	getmaxart := viper.GetString("common.listart")
	maxart, _ := strconv.Atoi(getmaxart)
	print(hide)
	if hide == "0" {
		db.Where("1=1").Order("id desc").Limit(maxart).Offset(maxart * (cpageid - 1)).Find(&articles)
	} else {
		db.Where("public = 0").Order("id desc").Limit(maxart).Offset(maxart * (cpageid - 1)).Find(&articles)
	}
	// 获取列表
	// 判断获取情况
	if len(articles) == 0 {
		response.NotFind(ctx, nil, "未找到已发布文章")
		return
	}

	for index, art := range articles {
		// 使用dto处理全部article
		items["art"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}
	response.Success(ctx, gin.H{
		"pageid": pageid,
		"arts":   items,
	}, "获取全部文章成功")
}

// Search 搜索用户名或文章名
func Search(ctx *gin.Context) {
	db := database.GetDB()

	key := ctx.PostForm("key")
	if len(key) < 0 {
		response.NotFind(ctx, nil, "什么都没找到~")
		return
	}

	// map处理全部articles
	var articles []model.Article
	items := make(map[string]model.ArticleDto)

	db.Where("title LIKE ?", "%"+key+"%").Where("public = 1").Find(&articles)

	// 判断获取情况
	if len(articles) == 0 {
		response.NotFind(ctx, nil, "什么都没找到")
		return
	}

	for index, art := range articles {
		// 使用dto处理全部article
		items["art"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}
	response.Success(ctx, gin.H{
		"arts": items,
	}, "搜索成功")
}

// GetActive 获取首页信息
func GetActive(ctx *gin.Context) {
	db := database.GetDB()
	// 获取参数

	maxart := viper.GetInt("common.maxart")

	navbarItem := viper.Get("homeinfo.navbar")
	newbarItem := viper.Get("homeinfo.newbar")
	footbarItem01 := viper.Get("homeinfo.footbar01")
	footbarItem02 := viper.Get("homeinfo.footbar02")
	footbarItem03 := viper.Get("homeinfo.footbar03")
	footbarItem04 := viper.Get("homeinfo.footbar04")
	footbarItem05 := viper.Get("homeinfo.footbar05")

	var articles01 []model.Article
	var articles02 []model.Article
	var articles03 []model.Article
	items01 := make(map[string]model.ArticleDto)
	items02 := make(map[string]model.ArticleDto)
	items03 := make(map[string]model.ArticleDto)

	db.Where("art_type = ?", "1").Where("public = 1").Order("id desc").Limit(maxart).Find(&articles01)
	db.Where("art_type = ?", "2").Where("public = 1").Order("id desc").Limit(maxart).Find(&articles02)
	db.Where("art_type = ?", "3").Where("public = 1").Order("id desc").Limit(maxart).Find(&articles03)

	for index, art := range articles01 {
		// 使用dto处理全部article
		items01["art"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}
	for index, art := range articles02 {
		// 使用dto处理全部article
		items02["art"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}
	for index, art := range articles03 {
		// 使用dto处理全部article
		items03["art"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}

	// 获取最近的文章
	response.Success(ctx, gin.H{
		"navbar":    navbarItem,
		"newbar":    newbarItem,
		"footbar01": footbarItem01,
		"footbar02": footbarItem02,
		"footbar03": footbarItem03,
		"footbar04": footbarItem04,
		"footbar05": footbarItem05,
		"art01":     items01,
		"art02":     items02,
		"art03":     items03,
	}, "获取首页信息成功")
}

// ModInfo 修改首页信息
func ModInfo(ctx *gin.Context) {
	// 获取参数
	key := ctx.PostForm("key")
	navid, err := strconv.Atoi(ctx.PostForm("navid"))
	if err != nil || !(navid > 0 && navid < 37) || len(key) == 0 {
		response.NotFind(ctx, nil, "请求参数错误")
		return
	}

	navbarItem := viper.GetStringSlice("homeinfo.navbar")
	newbarItem := viper.GetStringSlice("homeinfo.newbar")
	footbarItem01 := viper.GetStringSlice("homeinfo.footbar01")
	footbarItem02 := viper.GetStringSlice("homeinfo.footbar02")
	footbarItem03 := viper.GetStringSlice("homeinfo.footbar03")
	footbarItem04 := viper.GetStringSlice("homeinfo.footbar04")
	footbarItem05 := viper.GetStringSlice("homeinfo.footbar05")

	if navid < 9 {
		navbarItem[navid-1] = key
	} else if navid < 12 {
		newbarItem[navid-9] = key
	} else if navid < 17 {
		footbarItem01[navid-12] = key
	} else if navid < 22 {
		footbarItem02[navid-17] = key
	} else if navid < 27 {
		footbarItem03[navid-22] = key
	} else if navid < 32 {
		footbarItem04[navid-27] = key
	} else if navid < 37 {
		footbarItem05[navid-32] = key
	} else {
		response.NotFind(ctx, nil, "请求参数错误")
	}

	viper.Set("homeinfo.navbar", navbarItem)
	viper.Set("homeinfo.newbar", newbarItem)
	viper.Set("homeinfo.footbar01", footbarItem01)
	viper.Set("homeinfo.footbar02", footbarItem02)
	viper.Set("homeinfo.footbar03", footbarItem03)
	viper.Set("homeinfo.footbar04", footbarItem04)
	viper.Set("homeinfo.footbar05", footbarItem05)
	viper.WriteConfig()

	// 获取最近的文章
	response.Success(ctx, nil, "更改成功")
}
