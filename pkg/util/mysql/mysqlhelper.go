package mysql

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"acttos.com/avatar/pkg/setting"
	"acttos.com/avatar/pkg/util/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type MySQLHelper struct {
	Db *gorm.DB
}

var Helper *MySQLHelper
var once sync.Once

func InitMySQLHelper() {
	once.Do(func() {
		Helper = initMySQLDb()
	})
}

// Setup initialize the configuration instance
func initMySQLDb() *MySQLHelper {
	mySQLHelper := &MySQLHelper{}
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	host := setting.GetInstance().MysqlHost
	port := setting.GetInstance().MysqlPort
	dbName := setting.GetInstance().MysqlDbName
	user := setting.GetInstance().MysqlUser
	password := setting.GetInstance().MysqlPassword
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	logFile, _ := os.OpenFile(setting.GetInstance().GormLoggerFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	logLevel := gormLogger.Warn
	if setting.GetInstance().ServerMode == "debug" {
		logLevel = gormLogger.Info
	}
	newLogger := gormLogger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for gormLogger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Monitor.Error("Cannot open MySQL database")
		return nil
	}

	sqlDb, err := db.DB()
	if err == nil {
		sqlDb.SetMaxIdleConns(5)
		sqlDb.SetMaxOpenConns(10)
		sqlDb.SetConnMaxIdleTime(5 * time.Minute)
	}

	mySQLHelper.Db = db

	return mySQLHelper
}
