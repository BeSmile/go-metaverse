package live_helper

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-metaverse/models/live"
	"go-metaverse/models/live/constants"
	config2 "go-metaverse/tools/config"
	"os"
)

var (
	config   string
	platform string
	StartCmd = &cobra.Command{
		Use:  "lh",
		Long: `live-help`,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config/live.yml", "Start server with provided configuration file")

	StartCmd.PersistentFlags().StringVarP(&platform, "platform", "p", "douyu", "platform")
}

func run() {

	config2.ConfigLiveSetup(config)

	fmt.Println(config2.BiLiBiLiConfig.Domain)

	var client *live.ClientAdapter

	switch platform {
	case "douyu":
		client = live.NewClient(config2.DouyuConfig.Domain, config2.DouyuConfig.Port, constants.DouYu)
		break
	case "bilibili":
		client = live.NewClient(config2.BiLiBiLiConfig.Domain, config2.BiLiBiLiConfig.Port, constants.BiLiBiLi)
	}

	err := client.Connection()
	if err != nil {
		fmt.Println("连接弹幕失败")
		os.Exit(-1)
	}

	//client.JoinRoom("12235923")
	// b站房间号
	//client.JoinRoom("27848294")
	client.JoinRoom(23870434)
	client.HeartBeat()
	client.Watch()

	return
}
