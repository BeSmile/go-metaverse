package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var (
	cfgDatabase  *viper.Viper
	cfgApplition *viper.Viper
)

func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Read config file failed")
	}

	// 通过文件进行转换变量
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		fmt.Println("Parse config file failed")
	}

	// 配置数据库配置
	cfgDatabase = viper.Sub("settings.database")
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplition = viper.Sub("settings.application")
	ApplicationConfig = InitApplication(cfgApplition)

}

func SetConfig(configPath string, key string, value interface{}) error {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	return viper.WriteConfig()
}
