package handler

import (
	"practice-backend/model"
	"practice-backend/model/entity"

	"github.com/gofiber/fiber/v2"
)

func GetArticle(c *fiber.Ctx) error {
	parser := new(entity.ArticleQueryParser)

	if err := c.QueryParser(parser); err != nil {
		return err
	}

	articles, err := model.GetArticleData(*parser)
	if err != nil {
		return err
	}

	return c.JSON(articles)
}

func CreateArticle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": "This is my practice project",
	})
}
