package setting

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type App struct {
	PrefixUrl string

	RuntimeRootPath string

	LogSavePath        string
	LogSaveName        string
	LogFileExt         string
	TimeFormat         string
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
	if GetEnvStr("APP_ENV") == "TEST" {
		err:= godotenv.Overload()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	AppSetting.RuntimeRootPath = GetEnvStr("RUNTIME_ROOT_PATH")
	AppSetting.LogSavePath = GetEnvStr("LOG_SAVE_PATH")
	AppSetting.LogSaveName = GetEnvStr("LOG_SAVE_NAME")
	AppSetting.LogFileExt = GetEnvStr("LOG_FILE_EXT")
	AppSetting.TimeFormat = GetEnvStr("TIME_FORMAT")
	AppSetting.RunMode = GetEnvStr("RUN_MODE")

	PostgresSetting.Host = GetEnvStr("DB_HOST")
	PostgresSetting.Port = GetEnvStr("DB_PORT")
	PostgresSetting.User = GetEnvStr("DB_USER")
	PostgresSetting.Password = GetEnvStr("DB_PASSWORD")
	PostgresSetting.DatabaseName = GetEnvStr("DB_NAME")
}

func GetEnvStr(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal("Environment variable %s doesn't exist", key)
	}
	return v
}

func getenvInt(key string) int {
	s := GetEnvStr(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func getenvBool(key string) bool {
	s := GetEnvStr(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
