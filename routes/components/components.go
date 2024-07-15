package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetComponentsRoutes(app *fiber.App) {

	SetGlobalComponentRoutes(app)
	SetSummaryComponentRoutes(app)
}
