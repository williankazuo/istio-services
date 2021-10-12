package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello From service-c")
	})

	app.Listen(":3002")
}
