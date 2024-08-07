package pages

import (
	"eiyaro-htmx-explorer/models"

	"github.com/gofiber/fiber/v2"
)

func GetBlocks(c *fiber.Ctx) error {
	translation := c.Locals("lang").(models.Translation)

	return c.Render(
		"templates/index",
		fiber.Map{
			"Title": translation.M("pages").S("index"),
			"lang":  translation,
		},
		"templates/layout/global",
	)
}
