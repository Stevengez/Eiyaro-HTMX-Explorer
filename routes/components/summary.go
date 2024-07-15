package routes

import (
	webControllers "eiyaro-htmx-explorer/controllers/web"

	"github.com/gofiber/fiber/v2"
)

func SetSummaryComponentRoutes(app *fiber.App) {

	// BasePath for the lang selection
	summaryBaseRoute := app.Group("/summary")

	// all other dynamic calls
	summaryBaseRoute.Get("/lastblock", webControllers.GetLastBlock)
	summaryBaseRoute.Get("/lastBlocks", webControllers.GetLastBlocks)
	summaryBaseRoute.Get("/hashRate", webControllers.GetHashRate)
	summaryBaseRoute.Get("/difficulty", webControllers.GetDifficulty)
	summaryBaseRoute.Get("/pendings", webControllers.GetPendingCount)
}
