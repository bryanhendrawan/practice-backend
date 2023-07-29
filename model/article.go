package model

import (
	"practice-backend/database"
	"practice-backend/model/entity"
)

func GetArticles(query entity.ArticleQueryParser) ([]entity.Article, error) {
	var articles []entity.Article
	queryDB := database.DB

	if query.Author != "" {
		queryDB = queryDB.Where("author LIKE ?", "%"+query.Author+"%")
	}

	if query.Query != "" {
		queryDB = queryDB.Where("title LIKE ? OR body LIKE ?", "%"+query.Query+"%", "%"+query.Query+"%")
	}

	result := queryDB.Order("created desc").Find(&articles)

	return articles, result.Error
}

func CreateArticle(article entity.Article) (entity.Article, error) {
	result := database.DB.Create(&article)

	if result.Error != nil {
		return article, result.Error
	}

	return article, nil
}
