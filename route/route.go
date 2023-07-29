package route

import (
	"practice-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/articles", handler.GetArticle)
	r.Post("/articles", handler.CreateArticle)
}
