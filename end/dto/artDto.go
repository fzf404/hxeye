package dto

import (
	"App/model"
)

// ArticleDto 文章Dto
func ArticleDto(art model.Article) model.ArticleDto {
	timeTemplate := "2006-01-02 15:04"
	
	return model.ArticleDto{
		CreateAt: art.CreatedAt.Local().Format(timeTemplate),
		UpdateAt: art.UpdatedAt.Local().Format(timeTemplate),
		ArtID:    art.ID,
		UserID:   art.UserID,
		UserName: art.UserName,
		Title:    art.Title,
		ArtType:  art.ArtType,
		Content:  art.Content,
		Public:   art.Public,
	}
}

// ArticleInfoDto 文章Dto
func ArticleInfoDto(art model.Article) model.ArticleDto {
	timeTemplate := "2006-01-02 15:04"
	// 除Context
	return model.ArticleDto{
		CreateAt: art.CreatedAt.Local().Format(timeTemplate),
		UpdateAt: art.UpdatedAt.Local().Format(timeTemplate),
		ArtID:    art.ID,
		UserID:   art.UserID,
		UserName: art.UserName,
		Title:    art.Title,
		ArtType:  art.ArtType,
		Public:   art.Public,
	}
}
