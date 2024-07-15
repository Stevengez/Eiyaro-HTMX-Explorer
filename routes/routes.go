package routes

import (
	fs "eiyaro-htmx-explorer/fs"
	componentRoutes "eiyaro-htmx-explorer/routes/components"
	pageRoutes "eiyaro-htmx-explorer/routes/pages"

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

	componentRoutes.SetComponentsRoutes(app)
	pageRoutes.SetPagesRoutes(app)
}
