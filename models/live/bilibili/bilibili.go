package bilibili

/**
  https://github.com/lovelyyoshino/Bilibili-Live-API/blob/master/API.WebSocket.md
	https://github.com/Akegarasu/blivedm-go/blob/main/README.md
*/
import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/andybalholm/brotli"
	"go-metaverse/models/live/models"
	"go-metaverse/models/live/utils"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type AuthParams struct {
	Uid    int    `json:"uid"`
	Buvid  string `json:"buvid"`
	Roomid int    `json:"roomid"`
	// 默认3
	Protover int    `json:"protover"`
	Platform string `json:"platform"`
	// 默认2
	Type      int    `json:"type"`
	Key       string `json:"key"`
	ClientVer string `json:"clientver"`
}

type HostList struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	WssPort int    `json:"wss_port"`
	WsPort  int    `json:"ws_port"`
}

type RoomGroup struct {
	Group            string     `json:"group"`
	BusinessID       int        `json:"business_id"`
	Token            string     `json:"token"`
	RefreshRowFactor float32    `json:"refresh_row_factor"`
	RefreshRate      int        `json:"refresh_rate"`
	MaxDelay         int        `json:"max_delay"`
	HostList         []HostList `json:"host_list"`
}

type Response struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	TTL     int       `json:"ttl"`
	Data    RoomGroup `json:"data"`
}

// Encode 编码数据/**
/*
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-
| 					Packet Length 4 bytes			 |
| 	Header Length 2 bytes | Version Length 2 bytes   |
| 					Operation   4 bytes              |
| 					Sequence ID 4bytes               |
| 						Body ...                     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-

OP_HEARTBEAT	2	客户端发送的心跳包(30秒发送一次)
OP_HEARTBEAT_REPLY	3	服务器收到心跳包的回复
OP_SEND_SMS_REPLY	5	服务器推送的弹幕消息包
OP_AUTH	7	客户端发送的鉴权包(客户端发送的第一个包)
OP_AUTH_REPLY	8	服务器收到鉴权包后的回复

tcp传输数据为大端对齐
*/
func Encode(wsData string, msgType string) []byte {
	packetLen := int32(len(wsData) + 16)       // 需要计算body的长度加上header的固定长度, header的长度固定16
	lenByte := utils.IntToBytes(packetLen, "") // 4个
	//fmt.Println("packet", lenByte)
	headerByte := utils.Int16ToBytes(int16(16), "") // 2个
	//fmt.Println("header", headerByte)
	//如果Version=0，Body中就是实际发送的数据。
	//如果Version=2，Body中是经过压缩后的数据，请使用zlib解压，然后按照Proto协议去解析。
	versionByte := utils.Int16ToBytes(int16(1), "") // 2个
	//fmt.Println("ver", versionByte)
	var op int
	switch msgType {
	// OP_AUTH	7	客户端发送的鉴权包(客户端发送的第一个包)
	case "join":
		op = 7
		break
	// OP_HEARTBEAT	2	客户端发送的心跳包(30秒发送一次)
	case "heartbeat":
		op = 2
		break
	}
	opByte := utils.IntToBytes(int32(op), "") // 4个
	//fmt.Println("op", opByte)
	seqByte := utils.IntToBytes(int32(1), "") // 4个
	//fmt.Println("seq", seqByte)
	msgByte := []byte(wsData)
	//fmt.Println(msgByte, "msgByte")
	data := bytes.Join([][]byte{lenByte, headerByte, versionByte, opByte, seqByte, msgByte}, []byte(""))
	return data
}

type Client struct {
	Conn    net.Conn
	publish *models.Publisher
	roomId  int
	domain  string
	port    string
}

const (
	Plain = iota
	Popularity
	Zlib
	Brotli
)

func zlibParser(b []byte) ([]byte, error) {
	var rdBuf []byte
	zr, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	rdBuf, err = io.ReadAll(zr)
	return rdBuf, nil
}

func brotliParser(b []byte) ([]byte, error) {
	zr := brotli.NewReader(bytes.NewReader(b))
	rdBuf, err := io.ReadAll(zr)
	if err != nil {
		return nil, err
	}
	return rdBuf, nil
}

type Packet struct {
	Packet    []byte
	PacketLen uint32
	// 协议版本号 默认2 加密方式
	ProtoVersion uint16
	Header       uint16
	// 操作码
	Operation uint32
	// 保留字段，可以忽略
	SequenceId uint32
	Body       []byte
}

func (p *Packet) Decryption(data []byte) {
	fmt.Println(p.ProtoVersion, "Proto")
	switch p.ProtoVersion {
	case Popularity:
		fmt.Println("213445")
		fallthrough
	case Plain:
		msg := string(data[:len(data)-1])
		fmt.Println(msg, "msg4")
	case Zlib:
		z, err := zlibParser(data)
		if err != nil {
			log.Fatal("zlib error", err)
		}
		msg := string(data[:len(z)-1])
		fmt.Println(msg, "msg3")
	case Brotli:
		fmt.Println(data, "partBody")
		b, err := brotliParser(data)
		if err != nil {
			log.Fatal("Brotli error", err)
		}

		msg := string(data[:len(b)-1])
		fmt.Println(msg, "msg2")
	}
	msg := string(data[:len(data)-1])
	fmt.Println(msg, "msg1", len(data))
}

func NewPacket(proto uint16, operation uint32, body []byte) Packet {
	return Packet{
		ProtoVersion: proto,
		Operation:    operation,
		Body:         body,
	}
}

func NewPacketFromBytes(data []byte) Packet {
	packLen := binary.BigEndian.Uint32(data[0:4])
	pv := binary.BigEndian.Uint16(data[6:8])
	op := binary.BigEndian.Uint32(data[8:12])
	body := data[16:packLen]
	packet := NewPacket(pv, op, body)
	return packet
}

func Slice(data []byte) []Packet {
	var packets []Packet
	total := len(data)
	cursor := 0
	for cursor < total {
		packLen := int(binary.BigEndian.Uint32(data[cursor : cursor+4]))
		packets = append(packets, DecodePacket(data[cursor:cursor+packLen]))
		cursor += packLen
	}
	return packets
}

// DecodePacket Decode
func DecodePacket(data []byte) Packet {
	return NewPacketFromBytes(data)
}

// Parse 根据proto进行对数据的解密
func (p Packet) Parse() []Packet {

	switch p.ProtoVersion {
	case Popularity:
		fallthrough
		// 不处理
	case Plain:
		fmt.Println("body", string(p.Body))
		return []Packet{p}
	// Version=2时，zlib压缩后的body格式可能包含多个完整的proto包（可以理解为递归）。
	case Zlib:
		z, err := zlibParser(p.Body)
		if err != nil {
			log.Fatal("zlib error", err)
		}
		return Slice(z)
	case Brotli:
		b, err := brotliParser(p.Body)
		if err != nil {
			log.Fatal("Brotli error", err)
		}
		return Slice(b)
	default:
		log.Println("unknown Proto")
	}
	return nil
}

func (c *Client) Watch() {
	for {
		// 读取消息长度(4字节), b站的packLen包含了整个数据包
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
		// 根据首4个字节,获得body内容长度
		readLength := binary.BigEndian.Uint32(readLengthByte)
		// 只读取了前面4个字节,所以需要后面加上12个字节+readLengthByte
		packetBody := make([]byte, 12+readLength)
		_, err = c.Conn.Read(packetBody)
		if err != nil {
			if err == io.EOF {
				fmt.Println("packetBody出错")
			} else {
				fmt.Println("读取消息Body出错:", err)
			}
			break
		}
		intactBytes := bytes.Join([][]byte{readLengthByte, packetBody}, []byte(""))
		// 小于16个字节.表示没内容
		if len(intactBytes) < 16 {
			fmt.Println("长度不够跳过, 生成string", intactBytes, readLengthByte)
			continue
		}
		for _, pk := range DecodePacket(intactBytes).Parse() {
			fmt.Printf("协议版本号:%v, 操作符: %v\n", pk.ProtoVersion, pk.Operation)
			fmt.Println(string(pk.Body))
		}
	}
}

func (c *Client) HeartBeat() {
	for {
		err := c.SendData(Encode("[object Object]", "heartbeat"))
		if err != nil {
			log.Fatal("心跳失败" + err.Error())
		}
		log.Println("Send HeartBeat Success")
		time.Sleep(30 * time.Second)
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
	//go c.HeartBeat()

	return nil
}

func (c *Client) GetPublisher() *models.Publisher {
	return c.publish
}

func NewClient(domain string, port string) *Client {
	client := &Client{
		domain: domain,
		port:   port,
	}
	client.publish = models.NewPublisher()
	return client
}

// JoinRoom
// https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=27848294 获取key以及登录token
func (c *Client) JoinRoom(roomId int) error {
	resp, err := http.Get(fmt.Sprintf("https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=%s", roomId))
	if err != nil {
		log.Fatalln(err)
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalln(err)
	}
	// 获取token
	c.roomId = roomId
	auth := &AuthParams{
		Roomid:    c.roomId,
		Uid:       0,
		Protover:  3,
		Type:      2,
		Platform:  "web",
		Key:       response.Data.Token,
		ClientVer: "1.14.3",
	}
	data, err := json.Marshal(auth)
	if err != nil {
		fmt.Println("转换auth数据出错")
	}
	fmt.Println(string(data), "data")
	err = c.SendData(Encode(string(data), "join"))
	if err != nil {
		return err
	}
	return nil
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
