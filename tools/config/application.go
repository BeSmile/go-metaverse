package config

import "github.com/spf13/viper"

var ApplicationConfig = new(Application)

type Application struct {
	Domain       string
	Port         string
	Host         string
	IsHttps      bool
	Mode         string
	Name         string
	ReadTimeout  int
	WriteTimeout int
	JwtSecret    string
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		Domain:       cfg.GetString("domain"),
		Port:         portDefault(cfg),
		Host:         cfg.GetString("host"),
		IsHttps:      cfg.GetBool("ishttps"),
		Mode:         cfg.GetString("mode"),
		Name:         cfg.GetString("name"),
		ReadTimeout:  cfg.GetInt("readTimeout"),
		WriteTimeout: cfg.GetInt("writeTimeout"),
	}
}

func portDefault(cfg *viper.Viper) string {
	if cfg.GetString("port") == "" {
		return "8000"
	} else {
		return cfg.GetString("port")

	}
}
