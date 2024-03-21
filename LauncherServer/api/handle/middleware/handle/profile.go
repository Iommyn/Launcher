package handle

import (
	"LauncherServer/api/handle"
	"LauncherServer/internal/cache"
	"LauncherServer/internal/repository"
	"LauncherServer/internal/utils"
	"LauncherServer/internal/valider"
	"LauncherServer/internal/zaplogger"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"
	"time"
)

type validUsername struct {
	Username string `validate:"required,username,min=3,max=16"`
}

type profileResponse struct {
	UserID   int32  `json:"userID"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	Balance  int32  `json:"balance"`
}

func ProfileController(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	validData := &validUsername{
		Username: username,
	}

	zaplogger.Logger.Info("GetProfile: Валидация данных")

	if err := valider.Valider.Struct(validData); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, utils.ErrorValider(err.Field(), err.Tag()))
		}
		zaplogger.Logger.Error("GetProfile: Ошибка валидация данных")
		return handle.SendToClientError(ctx, strings.Join(errors, "; "), 400)
	}

	zaplogger.Logger.Info("GetProfile: Проверка есть ли в кэши данные")
	isProfileFound, err := cache.RedisCacheService.Exists("launcher_profile_" + username)

	if err != nil {
		zaplogger.Logger.Error("GetProfile: Ошибка получения данных с кэша", zap.Error(err))
		return handle.SendToClientError(ctx, "Ошибка получения данных с кэша", 500)
	}

	if !isProfileFound {
		zaplogger.Logger.Info("GetProfile: Получения с базы данных профиля")
		userID, user, isAdmin, err := repository.GetUserBaseProfileByUserName(username)

		if err != nil {
			if err == sql.ErrNoRows {
				zaplogger.Logger.Error("GetProfile: Ошибка получения с базы данных профиля, т.к. он отсутствует", zap.Error(err))
				return handle.SendToClientError(ctx, "Игрок не найден.", 404)
			}
			zaplogger.Logger.Error("GetProfile: Ошибка получения с базы данных профиля", zap.Error(err))
			return handle.SendToClientError(ctx, "Ошибка получения пользователя.", 500)
		}

		zaplogger.Logger.Info("GetProfile: Получения с базы данных профиля пользователя")
		balance, err := repository.GetUserBalanceByUserID(userID)
		if err != nil {
			zaplogger.Logger.Error("GetProfile: Ошибка получения с базы данных профиля пользователя", zap.Error(err))
			return handle.SendToClientError(ctx, "Ошибка получения профиля пользователя.", 500)
		}

		response := profileResponse{
			UserID:   userID,
			Username: user,
			IsAdmin:  isAdmin,
			Balance:  balance,
		}

		jsonData, err := json.Marshal(response)
		if err != nil {
			zaplogger.Logger.Error("GetProfile: Ошибка сериализации.", zap.Error(err))
			return handle.SendToClientError(ctx, "Ошибка сериализации.", 500)
		}

		err = cache.RedisCacheService.Set("launcher_profile_"+username, jsonData, time.Minute*15)

		if err != nil {
			zaplogger.Logger.Error("GetProfile: Ошибка кэширования", zap.Error(err))
			return handle.SendToClientError(ctx, "Ошибка кэширования.", 500)
		}

		zaplogger.Logger.Info("GetProfile: Профиль получен!", zap.Int32("userID", userID),
			zap.String("username", username), zap.Bool("isAdmin", isAdmin), zap.Int32("balance", balance))
		return ctx.Status(200).JSON(response)
	}

	dataJson, err := cache.RedisCacheService.Get("launcher_profile_" + username)

	if err != nil {
		zaplogger.Logger.Error("GetProfile: Ошибка получения профиля.", zap.Error(err))
		return handle.SendToClientError(ctx, "Ошибка получения профиля.", 500)
	}

	var profile profileResponse

	err = json.Unmarshal([]byte(dataJson), &profile)

	if err != nil {
		zaplogger.Logger.Error("GetProfile: Ошибка десериализации.", zap.Error(err))
		return handle.SendToClientError(ctx, "Ошибка десериализации.", 500)
	}

	zaplogger.Logger.Info("GetProfile: Профиль получен!", zap.Int32("userID", profile.UserID),
		zap.String("username", profile.Username), zap.Bool("isAdmin", profile.IsAdmin), zap.Int32("balance", profile.Balance))

	return ctx.Status(200).JSON(profile)
}
