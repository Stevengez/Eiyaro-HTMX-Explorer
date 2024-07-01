package routes

import (
	pagesControllers "eiyaro-htmx-explorer/controllers/pages"
	webControllers "eiyaro-htmx-explorer/controllers/web"
	fs "eiyaro-htmx-explorer/fs"
	middleware "eiyaro-htmx-explorer/middlewares/translation"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func SetRoutes(app *fiber.App) {

	// Serve Static Content
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(fs.StaticFS),
		PathPrefix: "static",
		Browse:     false,
	}))

	// all other dynamic calls
	app.Get("/langs", webControllers.GetLangs)

	// all other dynamic calls
	app.Get("/summary/lastblock", webControllers.GetLastBlock)
	app.Get("/summary/lastBlocks", webControllers.GetLastBlocks)
	app.Get("/summary/hashRate", webControllers.GetHashRate)
	app.Get("/summary/difficulty", webControllers.GetDifficulty)
	app.Get("/summary/pendings", webControllers.GetPendingCount)

	// BasePath for the lang selection
	langBaseRoute := app.Group("/:lang", middleware.Translation)

	// All Content Routes
	langBaseRoute.Get("/", pagesControllers.GetHome)

}
