package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/tap2joy/ChatService/utils"
)

const driverName = "postgres"

func syncDatabase(engine *xorm.Engine) error {
	return engine.Sync2(new(Channel), new(ChatLog))
}

//NewPGEngine PG orm 引擎
func NewPGEngine() *xorm.Engine {
	appConfig, err := utils.GetConfig("app.json")
	if err != nil {
		log.Printf("app config load failed %v\n", err)
		os.Exit(1)
	}

	dbName := appConfig.GetString("db_name")
	sslMode := appConfig.GetString("db_ssl_mode")
	dbHost := appConfig.GetString("db_host")
	dbPort := appConfig.GetInt("db_port")
	dbUser := appConfig.GetString("db_user")
	dbPwd := appConfig.GetString("db_password")
	dbMaxConnections := appConfig.GetInt("db_max_connections")
	dbIdleConnections := appConfig.GetInt("db_idle_connections")
	dbMaxLifeTime := appConfig.GetInt("db_max_life_time")

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbUser, dbPwd, dbHost, dbPort, dbName, sslMode)

	engine, err := xorm.NewEngine("postgres", connString)
	if err != nil {
		log.Fatalf("PG engine init error: %v", err)
	}

	engine.DatabaseTZ = time.Local
	engine.TZLocation = time.Local
	engine.SetMaxOpenConns(dbMaxConnections)
	engine.SetMaxIdleConns(dbIdleConnections)
	engine.SetConnMaxLifetime(time.Duration(dbMaxLifeTime) * time.Second)

	if err := engine.Ping(); err != nil {
		log.Fatalf("PG engine connect error: %v", err)
	}

	if err := syncDatabase(engine); err != nil {
		log.Fatalf("PG engine sync error: %v", err)
	}

	// if os.Getenv("APP_ENVIRONMENT") != "production" {
	// 	engine.ShowSQL()
	// }

	return engine
}

//Engine 全局数据库引擎
var Engine = NewPGEngine()
