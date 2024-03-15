package database

import (
	db "LauncherServer/db/sqlc"
	"LauncherServer/internal/utils"
	"LauncherServer/internal/zaplogger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var Store *db.Store

func InitDataBase(config utils.Config) *sql.DB {
	conn, err := sql.Open("mysql", config.DbConfig)
	if err != nil {
		zaplogger.Logger.Error("Неудалось подключиться к бд MySQL", zap.Error(err))
	}
	return conn
}
