package setting

import (
	"github.com/joho/godotenv"
	"log"
	"github.com/Celtcoste/server-graphql/utils"
)

type App struct {
	RunMode 		   string
}

var AppSetting = &App{}

type PostgreSQL struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

var PostgresSetting = &PostgreSQL{}

// Setup initialize the configuration instance
func Setup() {
	if utils.GetEnvStr("APP_ENV") == "TEST" {
		err:= godotenv.Overload()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	AppSetting.RunMode = utils.GetEnvStr("RUN_MODE")

	PostgresSetting.Host = utils.GetEnvStr("DB_HOST")
	PostgresSetting.Port = utils.GetEnvStr("DB_PORT")
	PostgresSetting.User = utils.GetEnvStr("DB_USER")
	PostgresSetting.Password = utils.GetEnvStr("DB_PASSWORD")
	PostgresSetting.DatabaseName = utils.GetEnvStr("DB_NAME")
}