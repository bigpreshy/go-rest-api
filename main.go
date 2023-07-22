package main

import "github.com/gofiber/fiber/v2"

func main() {
    server := fiber.New()

	app.Static("/", "./assets") 

    server.Get("/", func(request *fiber.Ctx) error {
        return request.SendString("Hello, World!gdg")
    })

	// server.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("value: " + c.Params("value"))
	// 	// => Get request with value: hello world
	// })

	server.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})



    server.Listen(":3000")
}