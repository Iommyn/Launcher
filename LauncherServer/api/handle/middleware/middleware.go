package middleware

import (
	"LauncherServer/api/handle"
	"LauncherServer/internal/utils"
	"LauncherServer/internal/zaplogger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func authMiddleware(config utils.Config) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			zaplogger.Logger.Error("Middleware: Ошибка на Frontend не указан заголовок authorization")
			return handle.SendToClientError(ctx, "Заголовок Authorization не указан", fiber.ErrUnauthorized.Code)
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			zaplogger.Logger.Error("Middleware: Ошибка на Frontend неверный формат заголовока authorization")
			return handle.SendToClientError(ctx, "Неверный формат заголовка авторизации", fiber.ErrUnauthorized.Code)
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			zaplogger.Logger.Error("Middleware: Ошибка на Frontend используется не тот протокол, нужен Bearer")
			return handle.SendToClientError(ctx, "Данный протокол авторизации не поддерживается", fiber.ErrUnauthorized.Code)
		}
		accessToken := fields[1]
		resp, err := http.Get(fmt.Sprintf("%s%s", config.AuthValidUrl, accessToken))
		if err != nil {
			zaplogger.Logger.Error("Middleware: Сервис авторизации недоступен ", zap.Error(err))
			return handle.SendToClientError(ctx, "Сервис авторизации недоступен", fiber.StatusInternalServerError)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			zaplogger.Logger.Error("Middleware: Токен истёк #1")
			return handle.SendToClientError(ctx, "Токен истёк", fiber.ErrUnauthorized.Code)
		}
		return ctx.Next()
	}
}
