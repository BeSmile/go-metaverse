package live_helper

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-metaverse/gui"
	"go-metaverse/models/live"
	"go-metaverse/models/live/constants"
	"go-metaverse/models/live/models"
	"go-metaverse/models/live/utils"
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

	dmApp := gui.NewClient()

	go func() {
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

		chatMsgHandler := func(_ string, message models.Message) {
			chatMsg, ok := message.(models.ChatMsgMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			fmt.Printf("%s(lv.%s): %s\n", chatMsg.Nn, chatMsg.Level, chatMsg.Txt)
			dmApp.NewChatMessage(chatMsg)
			utils.Speak(fmt.Sprintf("%s", chatMsg.Txt))
		}

		userEnterMsgHandler := func(_ string, message models.Message) {
			uEnter, ok := message.(models.UenterMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			uEnter.Txt = fmt.Sprintf("欢迎%s进入直播间", uEnter.Nn)
			dmApp.NewUEnterMessage(uEnter)
			//utils.Speak(uEnter.Txt)
		}
		spbcMsgHandler := func(_ string, message models.Message) {
			spbcMsg, ok := message.(models.SpbcMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			utils.Speak(fmt.Sprintf("感谢%s送出的%s", spbcMsg.Sn, spbcMsg.Gn))
		}
		dgbMsgHandler := func(_ string, message models.Message) {
			dgbMsg, ok := message.(models.DgbMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			utils.Speak(fmt.Sprintf("感谢%s送出的礼物", dgbMsg.Nn))
		}
		go client.AddEventListener(constants.UEnterType, userEnterMsgHandler)
		go client.AddEventListener(constants.ChatMsgType, chatMsgHandler)
		go client.AddEventListener(constants.SpbcType, spbcMsgHandler)
		go client.AddEventListener(constants.DgbType, dgbMsgHandler)

		//client.JoinRoom("12235923")
		// b站房间号
		//client.JoinRoom("27848294")
		client.JoinRoom(13233348)
		client.HeartBeat()
		client.Watch()
	}()
	dmApp.Init()
	return
}
