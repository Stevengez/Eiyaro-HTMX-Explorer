package routes

import (
	web "eiyaro-htmx-explorer/controllers/web/home"
	middleware "eiyaro-htmx-explorer/middlewares/translation"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {

	// Serve Static Content
	app.Static("/static", "./static")

	// BasePath for the lang selection
	langBaseRoute := app.Group("/:lang", middleware.Translation)

	// All Content Routes
	langBaseRoute.Get("/", web.GetHome)

}
