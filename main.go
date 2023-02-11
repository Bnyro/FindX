package main

import (
	"log"

	"github.com/bnyro/findx/config"
	"github.com/bnyro/findx/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

var Main string

func main() {
	config.Init()

	engine := html.New("./templates", ".html")

	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	app.Use(cors.New())

	app.Get("/", handlers.Home)
	app.Get("/search", handlers.Search)
	app.Get("/api", handlers.Api)
	app.Get("/ac", handlers.Suggest)

	app.Static("/static", "./static")
	app.Get("/proxy", handlers.Proxy)
	app.Get("/config", handlers.Config)
	app.Get("/opensearch.xml", handlers.Opensearch)

	log.Fatal(app.Listen(*config.Addr))
}
