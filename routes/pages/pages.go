package routes

import (
	pagesControllers "eiyaro-htmx-explorer/controllers/pages"
	middleware "eiyaro-htmx-explorer/middlewares/translation"

	"github.com/gofiber/fiber/v2"
)

func SetPagesRoutes(app *fiber.App) {

	// BasePath for the lang selection
	langBaseRoute := app.Group("/:lang", middleware.Translation)

	// All Content Routes
	//#Home
	langBaseRoute.Get("/", pagesControllers.GetHome)
	//#Block
	langBaseRoute.Get("/block/:height", pagesControllers.GetBlock)
}
