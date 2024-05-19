package middleware

import (
	"eiyaro-htmx-explorer/models"
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
)

func loadTranslations(lang string) (models.Translation, error) {
	filename := "translations/" + lang + ".json"
	file, err := os.ReadFile(filename)
	if err != nil {

		// If not english then fallback with english
		if lang != "en" {
			return loadTranslations("en")
		} else {
			return nil, err
		}
	}

	var translations models.Translation
	err = json.Unmarshal(file, &translations)
	if err != nil {
		return nil, err
	}

	return translations, nil
}

// Middleware to load the translation
func Translation(c *fiber.Ctx) error {
	// Get language and fallbacks to "en"
	lang := c.Params("lang")
	if lang == "" {
		lang = "en"
	}

	translations, err := loadTranslations(lang)
	if err != nil {
		return err
	}
	c.Locals("lang", translations)
	return c.Next()
}
