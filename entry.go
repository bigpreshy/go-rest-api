package main

// import "fmt"
import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/template/html/v2"

func main() {
	engine := html.New("./resources/views", ".html")

    server := fiber.New(fiber.Config{
        Views: engine,
})

	server.Static("/", "./assets") 

	

    server.Get("/", func(request *fiber.Ctx) error {
		return request.Render("index", fiber.Map{
            "Title": "Hello, World!",
        })
    })

	server.Post("/retry", func(c *fiber.Ctx) error {
		// Get first value from form field "name":
	 result :=  c.FormValue("first")
	 return c.SendString(result)


	//  return fmt.Println(c.FormValue("first"))
		// => "john" or "" if not exist
	  
		// ..
	  })

	server.Use(func(request *fiber.Ctx) error {
		return request.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})

	



    server.Listen(":3000")
}