package handle

import (
	"LauncherServer/internal/utils"
	"LauncherServer/internal/zaplogger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type launcherVersionResponse struct {
	Valid bool `json:"valid"`
}

// swagger:route GET /launcher/version/:version version
// version
// Frontend отсылает версию в backend где она проверяется.
// Версия состоит из числа, месяца и года. Например: 150324
// 15 - это число 03 - это месяц(Март) 24 - это год 2024
// responses:
//
//	200: launcherVersionResponse
func LauncherVersionController(ctx *fiber.Ctx, config utils.Config) error {
	version := ctx.Params("version")
	if len(version) == 0 || version != config.LauncherVersion {
		zaplogger.Logger.Error("CheckVersion: Проверка на версию не пройдена", zap.String("version_old", version),
			zap.String("version_new", config.LauncherVersion))
		response := launcherVersionResponse{
			Valid: false,
		}
		return ctx.Status(200).JSON(response)
	}
	response := launcherVersionResponse{
		Valid: true,
	}
	zaplogger.Logger.Info("CheckVersion: Проверка на версию пройдена", zap.String("version_old", version),
		zap.String("version_new", config.LauncherVersion))
	return ctx.Status(200).JSON(response)
}
