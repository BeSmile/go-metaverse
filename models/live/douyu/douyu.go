package douyu

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"go-metaverse/models/live/constants"
	"go-metaverse/models/live/models"
	"go-metaverse/models/live/utils"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// Encode
//    data = content + '\0'   # 数据部结尾为'\0'
//    data_length = len(data)+8  # 填入的消息长度=头部长度+数据部长度(包含结尾的'\0')
//    code = 689  # 消息类型(客户端发送给弹幕服务器的文本格式数据)
//    # 消息头部：消息长度(4字节)+消息类型(2字节)+加密字段(1字节,默认为0)+保留字段(1字节,默认为0)
//    head = data_length.to_bytes(4, 'little') + code.to_bytes(2,'little')+ (0).to_bytes(2,'little')
//    client.sendall(data_length.to_bytes(4, 'little') + head) # 消息长度出现两遍，二者相同
//    msg = (data).encode('utf-8')  # 使用utf-8编码 数据部分
//    client.sendall(bytes(msg))   # 发送数据部分
///**
func Encode(wsData string) []byte {
	dataLen := int32(len(wsData) + 9)
	msgByte := []byte(wsData)
	lenByte := utils.IntToBytes(dataLen, "little")
	sendByte := []byte{0xb1, 0x02, 0x00, 0x00}
	endByte := []byte{0x00} //
	data := bytes.Join([][]byte{lenByte, lenByte, sendByte, msgByte, endByte}, []byte(""))
	return data
}

//Decode:解析返回字节码为字符串
func Decode(msgBytes []byte) []string {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("出了错：%v msgBytes：%x msgStr: %s", err, msgBytes, string(msgBytes))
		}
	}()
	pos := 0
	var msg []string
	for pos < len(msgBytes) {
		contentLength := utils.BytesToInt(msgBytes[pos:pos+4], "little")
		content := bytes.NewBuffer(msgBytes[pos+12 : pos+3+contentLength]).String()
		msg = append(msg, content)
		pos = 4 + contentLength + pos
	}
	return msg
}

func ParseMsg(rawMsg string) utils.Response {
	res := make(utils.Response)
	attrs := strings.Split(rawMsg, "/")
	attrs = attrs[0 : len(attrs)-1]
	for _, attr := range attrs {
		if attr != "" {
			attr := strings.Replace(attr, "@s", "/", 1)
			attr = strings.Replace(attr, "@A", "@", 1)
			couple := strings.Split(attr, "@=")

			if len(couple) >= 2 {
				res[couple[0]] = couple[1]
			}
		}
	}
	return res
}

type Client struct {
	Conn    net.Conn
	publish *models.Publisher
	roomId  string
	domain  string
	port    string
}

func NewClient(domain string, port string) *Client {
	client := &Client{
		domain: domain,
		port:   port,
	}
	client.publish = models.NewPublisher()
	client.publish.AddMessageHandle(MessageHandle)
	return client
}

func (c *Client) SendData(bytesData []byte) error {
	_, err := c.Conn.Write(bytesData)
	if err != nil {
		if err == io.EOF {
			log.Panicln("Connection closed by remote side!")
		} else {
			log.Panicln("出错了")
		}
		return err
	}
	return nil
}

func (c *Client) HeartBeat() {
	for {
		//timestamp := time.Now().Unix()

		err := c.SendData(Encode("type@=mrkl/"))
		if err != nil {
			log.Fatal("心跳失败" + err.Error())
		}
		time.Sleep(45 * time.Second)
	}
}

func (c *Client) JoinRoom(roomId int32) error {
	c.roomId = strconv.Itoa(int(roomId))
	err := c.SendData(Encode(fmt.Sprintf("type@=loginreq/roomid@=%s/", c.roomId)))
	if err != nil {
		return err
	}
	groupErr := c.SendData(Encode(fmt.Sprintf("type@=joingroup/rid@=%s/gid@=-9999/", c.roomId)))
	if groupErr != nil {
		return groupErr
	}
	return nil
}

func (c *Client) Watch() {
	for {
		// 读取消息长度(4字节)
		readLengthByte := make([]byte, 4)
		_, err := c.Conn.Read(readLengthByte)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by remote side!")
			} else {
				fmt.Println("读取消息长度出错:", err)
			}
			break
		}
		readLength := binary.LittleEndian.Uint32(readLengthByte)
		//fmt.Println("接收消息长度:", readLength)
		// 超出2048字节的暂不处理
		if readLength > 2048 {
			continue
		}
		// 读取消息内容
		partBody := make([]byte, readLength)
		_, err = c.Conn.Read(partBody)
		if err != nil {
			fmt.Println("读取消息内容出错:", err)
			break
		}
		if len(partBody) < 9 {
			//fmt.Println("无效的消息长度:", len(partBody))
			continue
		}
		msg := string(partBody[8 : len(partBody)-1])
		fmt.Println(msg, "msg")
		c.publish.MessageHandle(msg)

		// 如果收到终止连接则退出循环
		if len(partBody) == 0 {
			break
		}
	}
}

func (c *Client) Connection() error {
	serverAddr := fmt.Sprintf("%s:%s", c.domain, c.port)
	fmt.Println("serverAddr", serverAddr)
	tcpServer, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		println("ResolveTCPAddr %s failed:", tcpServer, err.Error())
		return err
	}
	c.Conn, err = net.DialTimeout("tcp", serverAddr, 10*time.Second)

	if err != nil {
		fmt.Println("tcp connect timeout", err)
		return err
	}

	fmt.Println("连接弹幕服务器成功")
	go c.HeartBeat()

	return nil
}

func (c *Client) GetPublisher() *models.Publisher {
	return c.publish
}

func MessageHandle(p *models.Publisher, message string) {
	mappingData := ParseMsg(message)
	var msg models.Message

	mtype := mappingData["type"]
	if mtype == nil {
		return
	}
	msgType := utils.GetMsgTypeReflect(mappingData["type"])

	userInfo := models.UserInfo{}
	if err := mapstructure.Decode(mappingData, &userInfo); err != nil {
		fmt.Println(err, "error")
	}
	userInfo.Ic = utils.GetDomainAvatar(userInfo.Ic)

	base := models.Base{}
	if err := mapstructure.Decode(mappingData, &base); err != nil {
		fmt.Println(err, "error")
	}
	switch msgType {
	case constants.ChatMsgType:
		chatMsg := models.ChatMsgMessage{}
		if err := mapstructure.Decode(mappingData, &chatMsg); err != nil {
			fmt.Println("解析chatMsg", err)
		}
		chatMsg.UserInfo = userInfo
		chatMsg.Base = base
		fmt.Println(message)
		//fmt.Println("消息数据: ",mappingData)
		msg = models.Message(chatMsg)
		break

	case constants.UEnterType:
		uEnter := models.UenterMessage{}
		uEnter.UserInfo = userInfo
		uEnter.Base = base
		msg = models.Message(uEnter)
		break

	case constants.DgbType:
		dgbMessage := models.DgbMessage{}
		if err := mapstructure.Decode(mappingData, &dgbMessage); err != nil {
			fmt.Println("解析chatMsg", err)
		}
		dgbMessage.UserInfo = userInfo
		dgbMessage.Base = base
		msg = models.Message(dgbMessage)
		break
	case constants.SpbcType:
		spbcMessage := models.SpbcMessage{}
		if err := mapstructure.Decode(mappingData, &spbcMessage); err != nil {
			fmt.Println("解析chatMsg", err)
		}
		msg = models.Message(spbcMessage)
		break
	case constants.GgbbType:
	case constants.RankListType:
	case constants.RankUpType:
	case constants.SsdType:
		break
	case constants.NobleNumInfoType:
	case constants.FrankType:
		//message := Message{}
		//if err := mapstructure.Decode(mappingData, &message); err != nil {
		//	fmt.Println(err)
		//}
		////fmt.Println(message, "message")
		//p.Publish(message.GetType(), message)
		break
	default:
		//fmt.Println("todo接收消息内容:", mappingData)
		break
	}
	if msg != nil {

		p.Publish(msg.GetType(), msg)
	}
}
