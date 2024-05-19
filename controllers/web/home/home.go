package web

import (
	"eiyaro-htmx-explorer/models"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	translation := c.Locals("lang").(models.Translation)

	return c.Render(
		"index",
		fiber.Map{
			"Title": translation.M("pages").S("index"),
			"lang":  translation,
		},
		"layout/global",
	)
}
