package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-metaverse/database"
	"go-metaverse/global/orm"
	"go-metaverse/models/docker"
	"go-metaverse/router"
	"go-metaverse/tools"
	config2 "go-metaverse/tools/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// api/serverCmd represents the api/server command
var (
	config   string
	port     string
	mode     string
	StartCmd = &cobra.Command{
		Use:   "server",
		Short: "Short",
		Long:  `Long`,
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("api server called")
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func setup() {
	// 配置参数
	config2.ConfigSetup(config)
	// 连接数据库
	database.Setup()

	// 异步任务队列
}

func run() error {
	// 配置开发环境
	if mode != "" {
		config2.SetConfig(config, "settings.application.mode", mode)
	}

	r := router.InitRouter()

	defer func() {
		err := orm.DB.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	docker.InitCmdBackendEnv()

	//db.AutoMigrate(&User{})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	if port != "" {
		config2.SetConfig(config, "settings.application.port", port)
	}

	srv := &http.Server{
		Addr:    config2.ApplicationConfig.Host + ":" + config2.ApplicationConfig.Port,
		Handler: r,
	}
	go func() {

		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Printf("%s Server Run http://%s:%s/ \r\n",
		tools.GetCurrentTimeStr(),
		config2.ApplicationConfig.Host,
		config2.ApplicationConfig.Port)
	//fmt.Printf("%s Swagger URL http://%s:%s/swagger/index.html \r\n",
	//	tools.GetCurrentTimeStr(),
	//	config2.ApplicationConfig.Host,
	//	config2.ApplicationConfig.Port)
	//fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrntTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	//fmt.Printf("%s Shutdown Server ... \r\n", tools.GetCurrntTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		//logger.Fatal("Server Shutdown:", err)
	}
	//logger.Info("Server exiting")

	return nil
}
