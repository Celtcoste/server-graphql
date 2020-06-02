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

	if getenvStr("APP_ENV") == "TEST" {
		err:= godotenv.Overload()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	AppSetting.RuntimeRootPath = getenvStr("RUNTIME_ROOT_PATH")
	AppSetting.LogSavePath = getenvStr("LOG_SAVE_PATH")
	AppSetting.LogSaveName = getenvStr("LOG_SAVE_NAME")
	AppSetting.LogFileExt = getenvStr("LOG_FILE_EXT")
	AppSetting.TimeFormat = getenvStr("TIME_FORMAT")
	AppSetting.RunMode = getenvStr("RUN_MODE")

	PostgresSetting.Host = getenvStr("DB_HOST")
	PostgresSetting.Port = getenvStr("DB_PORT")
	PostgresSetting.User = getenvStr("DB_USER")
	PostgresSetting.Password = getenvStr("DB_PASSWORD")
	PostgresSetting.DatabaseName = getenvStr("DB_NAME")
}

func getenvStr(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal("Environment variable %s doesn't exist", key)
	}
	return v
}

func getenvInt(key string) int {
	s := getenvStr(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func getenvBool(key string) bool {
	s := getenvStr(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
