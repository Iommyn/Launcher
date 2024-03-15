package api

import (
	"LauncherServer/api/handle"
	"LauncherServer/internal/utils"
	"github.com/gofiber/fiber/v2"
	redoc "github.com/natebwangsut/fiber-redoc"
)

func RegisterRoutes(app *fiber.App, config utils.Config) {
	app.Static("/docs", "./docs")
	app.Get("/docs/*", redoc.Handler)

	launcherGroup := app.Group("/launcher")

	launcherGroup.Get("/:version", func(ctx *fiber.Ctx) error {
		return handle.LauncherVersionController(ctx, config)
	})
}
