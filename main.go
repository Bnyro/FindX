package main

import (
	"log"
	"os"
	"strings"

	"github.com/bnyro/findx/config"
	"github.com/bnyro/findx/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	config.Init()

	engine := html.New("./templates", ".html")

	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	app.Use(cors.New())

	app.Static("/static", "./static")
	app.Get("/", handlers.Home)
	app.Get("/search", handlers.Search)
	app.Get("/api", handlers.Api)
	app.Get("/ac", handlers.Suggest)
	app.Get("/proxy", handlers.Proxy)

	app.Get("/opensearch.xml", func(c *fiber.Ctx) error {
		bytes, _ := os.ReadFile("./opensearch.xml")
		descr := strings.Replace(string(bytes), "{{baseUrl}}", c.BaseURL(), -1)

		return c.Send([]byte(descr))
	})

	log.Fatal(app.Listen(*config.Addr))
}
