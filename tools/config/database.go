package config

import "github.com/spf13/viper"

var DatabaseConfig = new(Database)

type Database struct {
	Dbtype   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Dbtype:   cfg.GetString("dbtype"),
		Host:     cfg.GetString("host"),
		Port:     cfg.GetInt("port"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
		Name:     cfg.GetString("name"),
	}
}
