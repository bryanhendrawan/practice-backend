package main

import (
	"practice-backend/database"
	"practice-backend/database/migration"
	"practice-backend/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//init db
	database.DatabaseInit()

	//run migration table
	migration.RunMigration()

	app := fiber.New()

	//route
	route.RouteInit(app)

	app.Listen(":3000")
}
