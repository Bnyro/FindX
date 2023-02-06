package main

import (
	"flag"
	"log"

	"github.com/bnyrogo/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
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

	addr := flag.String("addr", ":8080", "address to listen on")
	flag.Parse()

	log.Fatal(app.Listen(*addr))
}