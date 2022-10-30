package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}

	return &dbConfig
}

func DBURL(config *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}
