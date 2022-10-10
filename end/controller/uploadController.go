package controller

import (
	"App/database"
	"App/dto"
	"App/model"
	"App/response"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// UploadImg 上传图片
func UploadImg(ctx *gin.Context) {
	// 图片验证
	file, err := ctx.FormFile("img")
	if err != nil {
		response.Fail(ctx, nil, "图片获取失败")
		return
	}

	path := viper.GetString("common.path")

	fileName := ctx.PostForm("imgname")
	print(fileName)
	if fileName != "" {
		file.Filename = fmt.Sprintf("%s%s", path, fileName)
		err = ctx.SaveUploadedFile(file, file.Filename)
		if err != nil {
			response.Fail(ctx, nil, "图片保存失败")
			return
		}
	} else {
		user, _ := ctx.Get("user")
		now := time.Now().Unix()
		file.Filename = fmt.Sprintf("%s%s%d%s", path, user.(model.User).Name, now, file.Filename)
		err = ctx.SaveUploadedFile(file, file.Filename)
		if err != nil {
			response.Fail(ctx, nil, "图片上传失败")
			return
		}
	}

	var imglist = []string{file.Filename[4:]}

	ctx.JSON(200, gin.H{
		"code": 200,
		"data": imglist,
		"msg":  "上传成功",
	})
}

// SetActive 设置首页信息
func SetActive(ctx *gin.Context) {
	db := database.GetDB()
	// 获取参数

	maxart := viper.GetInt("common.maxart")

	navbarItem := viper.Get("homeinfo.navbar")
	newbarItem := viper.Get("homeinfo.newbar")
	footbarItem := viper.Get("homeinfo.footbar")

	var articles01 []model.Article
	var articles02 []model.Article
	var articles03 []model.Article
	items01 := make(map[string]model.ArticleDto)
	items02 := make(map[string]model.ArticleDto)
	items03 := make(map[string]model.ArticleDto)

	db.Where("art_type = ?", "科技").Order("id desc").Limit(maxart).Find(&articles01)

	db.Where("art_type = ?", "校园").Order("id desc").Limit(maxart).Find(&articles02)

	db.Where("art_type = ?", "艺术").Order("id desc").Limit(maxart).Find(&articles03)

	for index, art := range articles01 {
		// 使用dto处理全部article
		items01["article"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}
	for index, art := range articles02 {
		// 使用dto处理全部article
		items02["article"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}
	for index, art := range articles03 {
		// 使用dto处理全部article
		items03["article"+strconv.Itoa(index+1)] = dto.ArticleInfoDto(art)
	}

	// 获取最近的文章
	response.Success(ctx, gin.H{
		"navbar":  navbarItem,
		"newbar":  newbarItem,
		"footbar": footbarItem,
		"art01":   items01,
		"art02":   items02,
		"art03":   items03,
	}, "获取全部文章成功")
}
