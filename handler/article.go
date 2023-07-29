package handler

import (
	"practice-backend/model"
	"practice-backend/model/entity"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetArticles(c *fiber.Ctx) error {
	parser := new(entity.ArticleQueryParser)

	if err := c.QueryParser(parser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed parse query param",
			"error":   err.Error(),
		})
	}

	articles, err := model.GetArticles(*parser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed get articles",
			"error":   err.Error(),
		})
	}

	return c.JSON(articles)
}

func CreateArticle(c *fiber.Ctx) error {
	body := new(entity.CreateArticle)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed parse request body",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Validate error",
			"error":   err.Error(),
		})
	}

	newArticle := entity.Article{
		Author:  body.Author,
		Title:   body.Title,
		Body:    body.Body,
		Created: time.Now(),
	}

	article, err := model.CreateArticle(newArticle)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create new article",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success create new article",
		"article": article,
	})
}
