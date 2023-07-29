package handler

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strings"
	"time"

	dbModel "practice-backend/database/model"
	"practice-backend/entity"
	redisModel "practice-backend/redis/model"

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

	lowercaseParser := entity.ArticleQueryParser{
		Query:  strings.ToLower(parser.Query),
		Author: strings.ToLower(parser.Author),
	}

	var articles []entity.Article
	hashKeyArticle := string(Hash(lowercaseParser))

	isSuccess, err := redisModel.Get(fmt.Sprintf(entity.KeyArticleByHash, hashKeyArticle), &articles)
	if err != nil {
		log.Println("Failed to get articles from redis")
	}

	if isSuccess {
		return c.JSON(articles)
	}

	articles, err = dbModel.GetArticles(*parser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed get articles",
			"error":   err.Error(),
		})
	}

	err = redisModel.Set(fmt.Sprintf(entity.KeyArticleByHash, hashKeyArticle), articles)
	if err != nil {
		log.Println("Failed to set articles to redis")
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

	article, err := dbModel.CreateArticle(newArticle)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create new article",
			"error":   err.Error(),
		})
	}

	err = redisModel.Set(fmt.Sprintf(entity.KeyArticleByID, article.ID), article)
	if err != nil {
		log.Printf("Failed to store article %d into redis", article.ID)
	}

	return c.JSON(fiber.Map{
		"message": "Success create new article",
		"article": article,
	})
}

func Hash(s entity.ArticleQueryParser) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	return b.Bytes()
}
