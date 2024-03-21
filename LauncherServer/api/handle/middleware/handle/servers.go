package handle

import (
	"LauncherServer/api/handle"
	"LauncherServer/internal/cache"
	"LauncherServer/internal/zaplogger"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ServerBaseConfig struct {
	DescriptionMods []string `json:"descriptionMods"`
	ServerVersion   string   `json:"serverVersion"`
}

type serverResponse struct {
	ServerName      string   `json:"serverName"`
	ServerVersion   string   `json:"serverVersion"`
	ServerOnline    int      `json:"serverOnline"`
	DescriptionMods []string `json:"descriptionMods"`
}

type serverConfigRunResponse struct {
	ServerName        string   `json:"serverName"`
	ServerVersion     string   `json:"serverVersion"`
	JavaVersion       string   `json:"javaVersion"`
	ServerJava        string   `json:"serverJava"`
	JvmArg            []string `json:"jvmArg"`
	ArgumentsOptimize []string `json:"argumentsOptimize"`
	NativesPath       string   `json:"nativesPath"`
	ClassPath         []string `json:"classPath"`
	MainClass         string   `json:"mainClass"`
	ClientArguments   []string `json:"clientArguments"`
	FullScreen        []string `json:"fullScreen"`
	AutoConnect       []string `json:"autoConnect"`
	DescriptionMods   []string `json:"descriptionMods"`
}

func ServerController(ctx *fiber.Ctx) error {
	files, err := ioutil.ReadDir("./profile_servers")
	if err != nil {
		zaplogger.Logger.Error("GetServers: Ошибка получение список профиля серверов", zap.Error(err))
		return handle.SendToClientError(ctx, "Ошибка сервера", fiber.StatusInternalServerError)
	}

	zaplogger.Logger.Info("GetServers: Получение информации о серверах")
	onlineServers := make(map[string]string)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			name := strings.TrimSuffix(file.Name(), ".json")
			key := "online_" + name
			online, err := cache.RedisCacheService.Get(key)
			if err != nil {
				onlineServers[name] = "0"
			} else {
				onlineServers[name] = online
			}
		}
	}

	zaplogger.Logger.Info("GetServers: Формируем запрос")
	var servers []serverResponse
	for serverName, serverOnline := range onlineServers {
		serverOnlineInt, err := strconv.Atoi(serverOnline)

		configFile, err := os.Open("./profile_servers/" + serverName + ".json")
		if err != nil {
			zaplogger.Logger.Error("Ошибка открытия файла конфигурации", zap.Error(err))
			continue
		}
		defer configFile.Close()

		var serverConfig ServerBaseConfig
		decoder := json.NewDecoder(configFile)
		err = decoder.Decode(&serverConfig)
		if err != nil {
			zaplogger.Logger.Error("Ошибка декодирования файла конфигурации", zap.Error(err))
			continue
		}

		servers = append(servers, serverResponse{
			ServerName:      serverName,
			ServerOnline:    serverOnlineInt,
			DescriptionMods: serverConfig.DescriptionMods,
			ServerVersion:   serverConfig.ServerVersion,
		})
	}

	return ctx.Status(200).JSON(servers)
}

func ServerRunController(ctx *fiber.Ctx) error {
	serverName := ctx.Params("servername")
	configFile, err := os.Open("./profile_servers/" + serverName + ".json")
	if err != nil {
		zaplogger.Logger.Error("Ошибка открытия файла конфигурации", zap.Error(err))
		return handle.SendToClientError(ctx, "Ошибка получение конфига", fiber.StatusInternalServerError)
	}
	defer configFile.Close()

	var serverConfig serverConfigRunResponse
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&serverConfig)
	if err != nil {
		zaplogger.Logger.Error("Ошибка декодирования файла конфигурации", zap.Error(err))
		return handle.SendToClientError(ctx, "Ошибка получение конфига", fiber.StatusInternalServerError)
	}
	return ctx.Status(200).JSON(serverConfig)
}
