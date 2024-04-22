package middleware

import (
	"LauncherServer/api/handle/middleware/handle"
	"LauncherServer/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutesMiddleware(app *fiber.App, config utils.Config) {
	launcherMiddle := app.Group("launcher/middle").Use(authMiddleware(config))

	launcherMiddle.Get("/profile/:username", handle.ProfileController)

	launcherMiddle.Get("/servers", handle.ServerController)
	launcherMiddle.Get("/server/run/:servername", handle.ServerRunController)
}
