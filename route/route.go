package route

import (
	"practice-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/articles", handler.GetArticles)
	r.Post("/article", handler.CreateArticle)
}
