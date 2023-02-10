package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/bnyro/findx/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	osaddr := flag.String("opensearch", "http://localhost:8080", "opensearch public url")

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
		descr := strings.Replace(string(bytes), "{{baseUrl}}", *osaddr, -1)

		return c.Send([]byte(descr))
	})

	flag.Parse()

	log.Fatal(app.Listen(*addr))
}
