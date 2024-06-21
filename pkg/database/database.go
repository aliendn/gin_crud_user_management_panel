package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=" + viper.GetString("database.host") +
		" port=" + viper.GetString("database.port") +
		" user=" + viper.GetString("database.user") +
		" dbname=" + viper.GetString("database.dbname") +
		" password=" + viper.GetString("database.password") +
		" sslmode=disable"
	var err error
	db, err = gorm.Open("postgres", dsn)
	return db, err
}

func GetDB() *gorm.DB {
	return db
}
