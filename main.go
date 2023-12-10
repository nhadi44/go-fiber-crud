package main

import (
	"github.com/go-fiber-crud/database"
	"github.com/go-fiber-crud/database/migration"
	"github.com/go-fiber-crud/helper"
	"github.com/go-fiber-crud/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {

	database.DatabaseInit()
	migration.MigrationInit()
	
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
	})

	app.Use(cors.New())
	app.Use(helmet.New())

	route.Router(app)

	err := app.Listen(":3001")

	helper.PanicHandler(err)

}
