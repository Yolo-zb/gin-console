package mysql

import (
	"github.com/Yolo-zb/gin-console/bootstrap"
	"github.com/Yolo-zb/gin-console/src/gorm"
	"github.com/spf13/viper"
)

var mysqlConfigs map[string]gorm.Conf

func InitGorm() {
	bootstrap.InitViper(".env", "env", 8898, true)
	mysqlConfigs = map[string]gorm.Conf{
		"localhost": {
			Host:     viper.GetString("DB_LOCALHOST_RW_HOST"),
			Port:     viper.GetInt("DB_LOCALHOST_RW_PORT"),
			Database: viper.GetString("DB_LOCALHOST_RW_DATABASE"),
			User:     viper.GetString("DB_LOCALHOST_RW_USERNAME"),
			Password: viper.GetString("DB_LOCALHOST_RW_PASSWORD"),
			MaxConn:  5,
			MaxOpen:  5,
		},
	}
	gorm.InitGormPool(mysqlConfigs)
}