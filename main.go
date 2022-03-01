package main

import (
	"log"
	"os"
	"time"

	"github.com/Reynadi531/phising-checker-api/routes"
	"github.com/Reynadi531/phising-checker-api/utils"

	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	task()
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Hours().Do(task)

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group("/api")
	routes.V1Router(api)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func task() {
	err := utils.DownloadLinks("links.txt")
	if err != nil {
		panic(err)
	}
}
