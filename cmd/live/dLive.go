package live

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-metaverse/gui"
	liveModel "go-metaverse/models/live"
	"go-metaverse/models/live/constants"
	"go-metaverse/models/live/models"
	"go-metaverse/models/live/utils"
	"os"
)

var (
	domain string
	port   string
	// 平台
	platform string
	StartCmd = &cobra.Command{
		Use:  "live",
		Long: `live`,
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

/**
douyu api
https://open.douyu.com/source/api/63
*/
func init() {
	StartCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "danmuproxy.douyu.com", "socket domain url")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8601", "domain port")
}

func run() {
	dmApp := gui.NewClient()

	go func() {
		client := liveModel.NewClient(domain, port, "douyu")

		err := client.Connection()

		if err != nil {
			fmt.Println("无法连接弹幕服务器 ", err)
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

		fmt.Println("Connect Success")
		go client.HeartBeat()
		// 橙子
		//client.JoinRoom("4549169")
		//client.JoinRoom("557171")
		// 冷狗
		//client.JoinRoom("3125893")
		//client.JoinRoom("8014243")
		client.JoinRoom(11578607)
		//client.JoinRoom("99999")
		// 乐乐直播间
		//client.JoinRoom("414194")
		// 刀 | 冷
		//client.JoinRoom("5103806")

		go client.Watch()
	}()
	dmApp.Init()

	return
}

func main() {
	run()
}
