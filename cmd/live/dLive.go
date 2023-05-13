package live

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-metaverse/gui"
	liveModel "go-metaverse/models/live"
	"net"
	"os"
	"time"
)

var (
	domain   string
	port     string
	StartCmd = &cobra.Command{
		Use:  "live",
		Long: `dyLive`,
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

//func Speak(str string) {
//	cmd := exec.Command("say", "--voice=Mei-Jia", str)
//
//	//cmd := exec.Command("say", "--voice=Sin-ji", str)
//	_, err := cmd.CombinedOutput()
//	// 检查命令是否执行成功
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

func run() {
	dmApp := gui.NewClient()

	go func() {
		client := liveModel.NewClient()
		serverAddr := fmt.Sprintf("%s:%s", domain, port)
		fmt.Println("serverAddr", serverAddr)
		tcpServer, err := net.ResolveTCPAddr("tcp", serverAddr)
		if err != nil {
			println("ResolveTCPAddr %s failed:", tcpServer, err.Error())
			os.Exit(1)
		}

		client.Conn, err = net.DialTimeout("tcp", serverAddr, 10*time.Second)

		if err != nil {
			fmt.Println("无法连接弹幕服务器 ", err)
		}

		chatMsgHandler := func(_ string, message liveModel.Message) {
			chatMsg, ok := message.(liveModel.ChatMsgMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			fmt.Printf("%s(lv.%s): %s\n", chatMsg.Nn, chatMsg.Level, chatMsg.Txt)
			dmApp.NewChatMessage(chatMsg)
			liveModel.Speak(fmt.Sprintf("%s", chatMsg.Txt))
		}

		userEnterMsgHandler := func(_ string, message liveModel.Message) {
			uEnter, ok := message.(liveModel.UenterMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			uEnter.Txt = fmt.Sprintf("欢迎%s进入直播间", uEnter.Nn)
			dmApp.NewUEnterMessage(uEnter)
			liveModel.Speak(uEnter.Txt)

		}
		spbcMsgHandler := func(_ string, message liveModel.Message) {
			spbcMsg, ok := message.(liveModel.SpbcMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			liveModel.Speak(fmt.Sprintf("感谢%s送出的%s", spbcMsg.Sn, spbcMsg.Gn))
		}
		dgbMsgHandler := func(_ string, message liveModel.Message) {
			dgbMsg, ok := message.(liveModel.DgbMessage)
			if !ok {
				fmt.Println("类型转换失败", ok)
			}
			liveModel.Speak(fmt.Sprintf("感谢%s送出的礼物", dgbMsg.Nn))
		}
		go client.AddEventListener(liveModel.UEnterType, userEnterMsgHandler)
		go client.AddEventListener(liveModel.ChatMsgType, chatMsgHandler)
		go client.AddEventListener(liveModel.SpbcType, spbcMsgHandler)
		go client.AddEventListener(liveModel.DgbType, dgbMsgHandler)

		fmt.Println("Connect Success")
		go client.HeartBeat()
		// 橙子
		//client.JoinRoom("4549169")
		//client.JoinRoom("557171")
		// 冷狗
		//client.JoinRoom("3125893")
		//client.JoinRoom("8014243")
		//client.JoinRoom("11578607")
		client.JoinRoom("99999")
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
