package dbconnection

import (
	"fmt"
	"project/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var conf = config.New()
var DB = getDB()

func getDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getStrConnection()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}

	return db
}

func getStrConnection() string {
	return fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		conf.HOST,
		conf.PORT,
		conf.USER,
		conf.PASSWORD,
		conf.DBNAME,
		conf.SSLMODE,
	)
}
