package handlers

import (
	"github.com/bnyro/findx/utilities"
	"github.com/bnyro/findx/web"
	"github.com/gofiber/fiber/v2"
)

func Proxy(c *fiber.Ctx) error {
	uri := c.Query("url")

	if utilities.IsBlank(uri) {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	body, contentType, err := web.Request(uri)

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	c.Set("Content-Type", string(contentType))
	return c.Send(body)
}
