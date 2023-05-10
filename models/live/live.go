package live

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Client struct {
	Conn    net.Conn
	publish *Publisher
	roomId  string
}

func NewClient() *Client {
	client := &Client{}
	client.publish = NewPublisher()
	return client
}

func (c *Client) AddEventListener(topic MsgType, consumer func(name string, msg Message)) {
	ch := make(chan Message, 1)
	c.publish.Subscribe(topic, ch)
	c.publish.AddEventListener(topic, consumer)
}

func (c *Client) SendData(msg string) error {
	_, err := c.Conn.Write(Encode(msg))
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

		err := c.SendData("type@=mrkl/")
		if err != nil {
			log.Fatal("心跳失败" + err.Error())
		}
		time.Sleep(45 * time.Second)
	}
}

func (c *Client) JoinRoom(roomId string) error {
	c.roomId = roomId
	c.SendData(fmt.Sprintf("type@=loginreq/roomid@=%s/", roomId))
	c.SendData(fmt.Sprintf("type@=joingroup/rid@=%s/gid@=-9999/", roomId))
	return nil
}

func (c *Client) Watch() error {

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

		c.publish.MessageHandle(msg)

		// 如果收到终止连接则退出循环
		if len(partBody) == 0 {
			break
		}
	}
	return nil
}
