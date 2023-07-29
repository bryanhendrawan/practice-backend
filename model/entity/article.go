package entity

import "time"

type Article struct {
	ID      int       `json:"id" gorm:"primaryKey"`
	Author  string    `json:"author"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

type ArticleQueryParser struct {
	Query  string `query:"query"`
	Author string `query:"author"`
}

type CreateArticle struct {
	Author string `json:"author" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
}
