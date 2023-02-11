package handlers

import (
	"strings"

	_ "embed"

	"github.com/bnyro/findx/config"
	"github.com/gofiber/fiber/v2"
)

//go:embed opensearch.xml
var opensearchBody string

func Config(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"proxy":     *config.Proxy,
		"redirects": config.Redirects,
	})
}

func Opensearch(c *fiber.Ctx) error {
	descr := strings.Replace(opensearchBody, "{{baseUrl}}", c.BaseURL(), -1)
	return c.Send([]byte(descr))
}
