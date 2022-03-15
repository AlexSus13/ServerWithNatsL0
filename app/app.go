package app

import (
	"ServerWithNatsL0/config"
	"ServerWithNatsL0/model"

	"github.com/sirupsen/logrus"

	"database/sql"
)

type App struct {
	Db       *sql.DB
	MyLogger *logrus.Logger
	Config   *config.Conf
	ChanData chan model.OrdersData
	CacheMap map[string]model.OrdersData
}

func NewApp(db *sql.DB, MyLogger *logrus.Logger, Conf *config.Conf, Chan chan model.OrdersData, Cache map[string]model.OrdersData) *App {
	return &App{
		Db:       db,
		MyLogger: MyLogger,
		Config:   Conf,
		ChanData: Chan,
		CacheMap: Cache,
	}
}
