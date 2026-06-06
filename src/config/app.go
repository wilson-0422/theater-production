package config

type AppConfigStruct struct {
	AppName       string
	Port          string
	SessionSecret string
	DBPath        string
}

var AppConfig = AppConfigStruct{
	AppName:       "剧目排演管理系统",
	Port:          "8080",
	SessionSecret: "theater-production-secret-key-2024",
	DBPath:        "theater.db",
}
