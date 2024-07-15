package routes

import (
	webControllers "eiyaro-htmx-explorer/controllers/web"

	"github.com/gofiber/fiber/v2"
)

func SetGlobalComponentRoutes(app *fiber.App) {
	// Langs List Menu Component
	app.Get("/langs", webControllers.GetLangs)
}
