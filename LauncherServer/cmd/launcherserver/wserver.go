package launcherserver

import (
	"LauncherServer/api"
	db "LauncherServer/db/sqlc"
	"LauncherServer/internal/cache"
	"LauncherServer/internal/database"
	"LauncherServer/internal/utils"
	"LauncherServer/internal/valider"
	"LauncherServer/internal/zaplogger"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"log"
	"time"
)

type WebServer struct {
	App *fiber.App
}

func settingsCors(app *fiber.App, config utils.Config) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.DomainAllow,
		AllowMethods:     config.MethodsAllow,
		AllowHeaders:     config.HeadersAllow,
		AllowCredentials: config.IsCredentials,
	}))
}

func settingsApp(config utils.Config) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:               config.AppName,
		BodyLimit:             1024 * 1024,
		EnablePrintRoutes:     true,
		IdleTimeout:           10 * time.Second,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ReadTimeout:           10 * time.Second,
		WriteTimeout:          10 * time.Second,
		ProxyHeader:           fiber.HeaderXForwardedFor,
		ReadBufferSize:        1024 * 4,
		WriteBufferSize:       1024 * 4,
		ReduceMemoryUsage:     true,
		DisableStartupMessage: false,
	})
}

func (web *WebServer) StartLauncherService() error {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println(err)
	}

	zaplogger.InitializeLogger(config.SaveLogFile, config.MaxSizeFile, config.MaxBackupFile, config.MaxAgeFile)
	zaplogger.Logger.Info("Конфигурация загружена успешно!")
	zaplogger.Logger.Info("Логгер инциализирован!")

	app := settingsApp(config)
	zaplogger.Logger.Info("Настройка Fiber успешно прошла")

	app.Use(logger.New())
	zaplogger.Logger.Info("Инциализация логгера для Fiber")

	settingsCors(app, config)
	zaplogger.Logger.Info("Настройка прав доступа к доменам CORS")

	web.App = app
	zaplogger.Logger.Info("Инциализация структуры")

	valider.InitValider()
	zaplogger.Logger.Info("Валидация успешно инциализирована")

	database.Store = db.NewStore(database.InitDataBase(config))
	zaplogger.Logger.Info("Подключение к базе данных MySQL")

	cache.InitCacheRedis(config)
	zaplogger.Logger.Info("Соединение с кэшированием Redis")

	api.RegisterRoutes(app, config)
	zaplogger.Logger.Info("Регистрация роутов")

	if config.Ssl {
		zaplogger.Logger.Info("Запуск службы лаунчера используя SSL", zap.String("IP", config.ServerHttpAddress))
		return app.ListenTLS(config.ServerHttpAddress, config.SslCrt, config.SslKey)
	}

	zaplogger.Logger.Info("Запуск службы лаунчера без SSL", zap.String("IP", config.ServerHttpAddress))
	return app.Listen(config.ServerHttpAddress)
}
