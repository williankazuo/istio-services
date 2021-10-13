package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Variables:")
	fmt.Println("URL_SERVICE_A: " + os.Getenv("URL_SERVICE_A"))
	fmt.Println("URL_SERVICE_B: " + os.Getenv("URL_SERVICE_B"))

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello From service-c")
	})

	app.Get("/hello/a", func(c *fiber.Ctx) error {
		urlA := os.Getenv("URL_SERVICE_A")
		resp, err := http.Get(urlA + "/hello")
		if err != nil {
			return c.SendString(err.Error())
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendString(string(bodyBytes))
	})

	app.Get("/hello/b", func(c *fiber.Ctx) error {
		urlB := os.Getenv("URL_SERVICE_B")
		resp, err := http.Get(urlB + "/hello")
		if err != nil {
			return c.SendString(err.Error())
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendString(string(bodyBytes))
	})

	app.Listen(":3002")
}
