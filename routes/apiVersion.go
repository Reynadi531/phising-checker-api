package routes

import (
	v1 "github.com/Reynadi531/phising-checker-api/controller/v1"
	"github.com/gofiber/fiber/v2"
)

func V1Router(r fiber.Router) {
	v1route := r.Group("/v1")

	v1route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	v1route.Get("/check/", v1.CheckController)
}
