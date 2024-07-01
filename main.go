package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	fs "eiyaro-htmx-explorer/fs"
	"eiyaro-htmx-explorer/models"
	"eiyaro-htmx-explorer/routes"
	"eiyaro-htmx-explorer/service/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func errorHandler(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		if requestErr, ok := err.(*models.RequestError); ok {
			return c.Status(requestErr.Status).JSON(fiber.Map{
				"error": requestErr.Message,
			})
		} else if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
			// If the erros is of class json.UnmarshalTypeError
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bad Payload Format: " + jsonErr.Error(),
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

	}
	return nil
}

func main() {
	// Load env variables from .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env:", err)
	}

	// Initialize Fiber Wrapper for HTML Template Engine
	engine := html.NewFileSystem(http.FS(fs.TemplatesFS), ".html")
	engine.AddFunc("shorter", template.Shorter)

	// Reaload template files on each render
	enableReaload := os.Getenv("HOT_TEMPLATE_RELOAD")
	if enableReaload == "true" {
		engine.Reload(true)
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Apply cors policies :TODO:
	app.Use(cors.New())

	// Error Handler
	app.Use(errorHandler)

	// Set Server Routes
	routes.SetRoutes(app)

	serverHost := os.Getenv("SERVER_HOST") // IP of the interface where the server will serve
	serverPort := os.Getenv("SERVER_PORT") // Port for the server, if not present a random port will be used
	err := app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))

	if err != nil {
		panic(err)
	}
}
