package router

import (
	"github.com/kiliczsh/gondit/handler"
	"github.com/kiliczsh/gondit/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Scan
	scan := api.Group("/scan")
	scan.Get("/", handler.GetAllScans)
	scan.Get("/:id", handler.GetScan)
	scan.Get("/:id/test", handler.Test)
	scan.Post("/", middleware.Protected(), handler.CreateScan)
	scan.Delete("/:id", middleware.Protected(), handler.DeleteScan)
}