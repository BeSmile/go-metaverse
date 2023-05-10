package migrate

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-metaverse/database"
	orm "go-metaverse/global/orm"
	"go-metaverse/models/gorm"
	config2 "go-metaverse/tools/config"
)

var (
	mode     string
	config   string
	StartCmd = &cobra.Command{
		Use:   "migrate",
		Short: "mg",
		Long:  `migrate`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("migrate/server called", args, config)
			run()
		},
	}
)

// 注册参数命令 config 以及 mode命令
func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "dev | test | prod")
}

func run() {
	fmt.Printf("migrate %s\n\r", config)

	// 读取配置
	config2.ConfigSetup(config)

	// 初始化数据库连接
	database.Setup()

	// 数据库迁移
	_ = migrateModel()

	fmt.Println("数据库结构初始化成功！")
}

func migrateModel() error {
	// 作用?
	if config2.DatabaseConfig.Dbtype == "mysql" {
		orm.DB = orm.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return gorm.AutoMigrate(orm.DB)
}
