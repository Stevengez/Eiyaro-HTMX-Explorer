package web

import (
	"eiyaro-htmx-explorer/fs"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetLangs(c *fiber.Ctx) error {
	//translation := c.Locals("lang").(models.Translation)

	files, err := fs.LocalesFS.ReadDir("locales")

	if err != nil {
		fmt.Println("Failed to load locales")
	}

	langList := make([]string, len(files))

	for i, file := range files {
		langList[i] = strings.TrimSuffix(file.Name(), ".json")
		langList[i] = strings.ToUpper(langList[i])
	}

	return c.Render(
		"templates/layout/menu/langitem",
		fiber.Map{
			"Options": langList,
		},
	)
}
