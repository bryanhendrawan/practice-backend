package handler

import "github.com/gofiber/fiber/v2"

func GetArticle(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "This is my practice project",
	})
}
