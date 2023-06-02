package config

import (
	"github.com/spf13/viper"
)

type DouYu struct {
	Domain string
	Port   string
}

type BiLiBiLi struct {
	Domain  string
	Port    string
	WsPort  string
	WssPort string
}

type Live struct {
	DouYu
	BiLiBiLi
}

var DouyuConfig = new(DouYu)
var BiLiBiLiConfig = new(BiLiBiLi)

func InitDouYuConfig(cfg *viper.Viper) *DouYu {
	return &DouYu{
		Port:   cfg.GetString("port"),
		Domain: cfg.GetString("domain"),
	}
}

func InitBiliBiLiConfig(cfg *viper.Viper) *BiLiBiLi {
	return &BiLiBiLi{
		Port:    cfg.GetString("port"),
		Domain:  cfg.GetString("domain"),
		WssPort: cfg.GetString("wss_port"),
		WsPort:  cfg.GetString("ws_port"),
	}
}
