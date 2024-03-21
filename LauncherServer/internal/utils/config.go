package utils

import "github.com/spf13/viper"

type Config struct {
	AppName           string `mapstructure:"APP_NAME"`
	DomainAllow       string `mapstructure:"DOMAIN_ALLOW"`
	MethodsAllow      string `mapstructure:"ALLOW_METHODS"`
	HeadersAllow      string `mapstructure:"ALLOW_HEADERS"`
	IsCredentials     bool   `mapstructure:"ALLOW_CREDENTIALS"`
	DbConfig          string `mapstructure:"DB_CONFIG"`
	ServerHttpAddress string `mapstructure:"SERVER_HTTP_ADDRESS"`
	RedisAddress      string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword     string `mapstructure:"REDIS_PASSWORD"`
	RedisDB           int    `mapstructure:"REDIS_DB"`
	Ssl               bool   `mapstructure:"SSL"`
	SslKey            string `mapstructure:"SSL_KEY"`
	SslCrt            string `mapstructure:"SSL_CRT"`
	SaveLogFile       string `mapstructure:"SAVE_LOG_FILE"`
	MaxSizeFile       int    `mapstructure:"LOG_SIZE_FILE"`
	MaxBackupFile     int    `mapstructure:"LOG_BACKUP_FILE_COUNT"`
	MaxAgeFile        int    `mapstructure:"LOG_MAX_AGE"`
	LauncherVersion   string `mapstructure:"LAUNCHER_VERSION"`
	AuthValidUrl      string `mapstructure:"LAUNCHER_VERSION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
