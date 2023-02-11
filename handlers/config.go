package handlers

import (
	"os"
	"strings"

	"github.com/bnyro/findx/config"
	"github.com/gofiber/fiber/v2"
)

func Config(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"proxy":     *config.Proxy,
		"redirects": config.Redirects,
	})
}

func Opensearch(c *fiber.Ctx) error {
	bytes, _ := os.ReadFile("./opensearch.xml")
	descr := strings.Replace(string(bytes), "{{baseUrl}}", c.BaseURL(), -1)
	return c.Send([]byte(descr))
}
