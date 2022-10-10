package model

import "gorm.io/gorm"

// Article 文章数据
type Article struct {
	gorm.Model
	UserID   uint   // 用户id
	UserName string // 用户名
	Title    string // 文章标题
	ArtType  int    // 文章类型
	Content  string // 文章内容
	Public   bool   // 加精
}

// ArticleDto 文章数据传输
type ArticleDto struct {
	CreateAt string `json:"create"`   // 创建时间
	UpdateAt string `json:"update"`   // 修改事件
	ArtID    uint   `json:"artid"`    // 文章id
	UserID   uint   `json:"userid"`   // 用户id
	UserName string `json:"username"` // 用户名
	Title    string `json:"title"`    // 文章标题
	ArtType  int    `json:"arttype"`  // 文章类型
	Content  string `json:"content"`  // 文章内容
	Public   bool   `json:"public"`   // 加精
}
