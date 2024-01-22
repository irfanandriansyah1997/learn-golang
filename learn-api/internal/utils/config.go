package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"API_PORT"`
	MysqlRootPassword string `mapstructure:"MYSQL_ROOT_PASSWORD"`
	MysqlDatabase     string `mapstructure:"MYSQL_DATABASE"`
	MysqlUser         string `mapstructure:"MYSQL_USER"`
	MysqlPassword     string `mapstructure:"MYSQL_PASSWORD"`
}

var AppConfig *Config

func LoadAppConfig(extensionName, fileName, pathName string) {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(pathName)
	viper.SetConfigName(fileName)
	viper.SetConfigType(extensionName)

	err := viper.ReadInConfig()
	PanicIfError(err)

	err = viper.Unmarshal(&AppConfig)

}
