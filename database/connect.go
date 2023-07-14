package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	//* Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=true", Config("DB_USER_DEV"), Config("DB_PASSWORD_DEV"), Config("DB_PORT_DEV"), Config("DB_NAME"))
	// dsn := "root:mysecretpassword@tcp(localhost:3306)/article_db?parseTime=true"
	//* Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Info)),
		// Logger: logger.Default.LogMode(logger.Silent),
		// SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
}
