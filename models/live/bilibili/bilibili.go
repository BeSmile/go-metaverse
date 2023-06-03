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
	"github.com/dlclark/regexp2"
	"go-metaverse/models/live/bilibili/message"
	"go-metaverse/models/live/constants"
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
	Type int    `json:"type"`
	Key  string `json:"key"`
	//ClientVer string `json:"clientver"`
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
	packetLen := int32(len(wsData) + 16)            // 需要计算body的长度加上header的固定长度, header的长度固定16
	lenByte := utils.IntToBytes(packetLen, "")      // 4个
	headerByte := utils.Int16ToBytes(int16(16), "") // 2个
	//如果Version=0，Body中就是实际发送的数据。
	//如果Version=2，Body中是经过压缩后的数据，请使用zlib解压，然后按照Proto协议去解析。
	versionByte := utils.Int16ToBytes(int16(1), "") // 2个
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
	seqByte := utils.IntToBytes(int32(1), "") // 4个
	msgByte := []byte(wsData)
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

// Slice 一条数据可能存在多个消息体,所以分割下多条
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
		if len(intactBytes) < 16 || len(intactBytes) > 20480 {
			//fmt.Println("长度不够跳过, 生成string", intactBytes, readLengthByte)
			continue
		}
		for _, pk := range DecodePacket(intactBytes).Parse() {

			c.Handle(pk)
		}
	}
}

const (
	_ = iota
	_
	HEART_BEAT
	HEART_BEAT_RESPONSE
	_
	// NOTICE 弹幕、广播等全部信息
	NOTICE
	_
	// JOIN_ROOM 进房
	JOIN_ROOM
	// WELCOME 进房通知
	ENTER_ROOM_RESPONSE
)

const (
	// 弹幕消息
	DANMU_MSG = "DANMU_MSG"
	// ENTRY_EFFECT {"cmd":"ENTRY_EFFECT","data":{"id":135,"uid":1970739762,"target_id":381245965,"mock_effect":0,"face":"https://i0.hdslb.com/bfs/face/member/noface.jpg","privilege_type":0,"copy_writing":"欢迎 \u003c%bili_56...%\u003e 直播间","copy_color":"#000000","highlight_color":"#FFF100","priority":1,"basemap_url":"https://i0.hdslb.com/bfs/live/mlive/da6933ea70f31c4df63f4b68b735891284888357.png","show_avatar":1,"effective_time":1,"web_basemap_url":"https:0.hdslb.com/bfs/live/mlive/da6933ea70f31c4df63f4b68b735891284888357.png","web_effective_time":2,"web_effect_close":0,"web_close_time":900,"business":3,"copy_writing_v2":"欢迎 \u003c^icon^\u003e \u003c%bili_5…%\u003e 进入直播间","st":[1],"max_delay_time":7,"trigger_time":1685699415328175480,"identities":22,"effect_silent_time":0,"effective_time_new":0,"web_dynamic_url_webp":"","web_dynamic_url_apng":"","mobile_dynamic_url_webp":""}}
	ENTRY_EFFECT = "ENTRY_EFFECT"
	// 欢迎房管
	WELCOME_GUARD = "WELCOME_GUARD"
	// 欢迎xxx进入房间
	WELCOME = "WELCOME_GUARD"
	// 在某一用户第一次进入直播间时会发送一条消息，用户退出直播间约十分钟之后，再次进入才会再次发送消息
	INTERACT_WORD = "INTERACT_WORD"
	// 多少人看过
	WATCHED_CHANGE = "WATCHED_CHANGE"

	// sc
	SUPER_CHAT_MESSAGE     = "SUPER_CHAT_MESSAGE"
	SUPER_CHAT_MESSAGE_JPN = "SUPER_CHAT_MESSAGE_JPN"

	// 高能榜数量
	ONLINE_RANK_COUNT    = "ONLINE_RANK_COUNT"
	ONLINE_RANK_COUNT_V2 = "ONLINE_RANK_COUNT_V2"
	STOP_LIVE_ROOM_LIST  = "STOP_LIVE_ROOM_LIST"
	// 点赞
	LINK_INFO_V3_UPDATE = "LINK_INFO_V3_UPDATE"
	LIKE_INFO_V3_UPDATE = "LIKE_INFO_V3_UPDATE"
)

const (
	// 投喂礼物
	SEND_GIFT = "SEND_GIFT"
	// 连击礼物
	// COMBO_SEND {"cmd":"COMBO_SEND","data":{"action":"投喂","batch_combo_id":"batch:gift:combo_id:4668669:381245965:31039:1685699397.2412","batch_combo_num":10,"combo_id":"gift:combo_id:4668669:381245965:31039:1685699397.2402","combo_m":10,"combo_total_coin":1000,"dmscore":64,"gift_id":31039,"gift_name":"牛哇牛哇","gift_num":0,"is_join_receiver":false,"is_naming":false,"is_show":1,"medal_info":{"anchor_roomid":0,"anchor_uname":"","guard_level":0,"icon_id":0,"ighted":0,"medal_color":1725515,"medal_color_border":12632256,"medal_color_end":12632256,"medal_color_start":12632256,"medal_level":22,"medal_name":"PuFF","special":"","target_id":1526101},"name_color":"","r_uname":"泰莉亚子","ree_user_info":{"uid":381245965,"uname":"泰莉亚子"},"ruid":381245965,"send_master":null,"total_num":10,"uid":4668669,"uname":"挽留下落"}}
	COMBO_SEND = "COMBO_SEND"
)

const (
	// 上舰长
	GUARD_BUY = "GUARD_BUY"
	// 续费了舰长
	USER_TOAST_MSG = "USER_TOAST_MSG"
	// 在本房间续费了舰长
	NOTICE_MSG = "NOTICE_MSG"
)

const (
	// 小时榜变动
	ACTIVITY_BANNER_UPDATE_V2 = "ACTIVITY_BANNER_UPDATE_V2"
	ONLINE_RANK_V2            = "ONLINE_RANK_V2"
)

const (
	// 粉丝关注变动  {"cmd":"ROOM_REAL_TIME_MESSAGE_UPDATE","data":{"roomid":22091661,"fans":116755,"red_notice":-1,"fans_club":42}}
	ROOM_REAL_TIME_MESSAGE_UPDATE = "ROOM_REAL_TIME_MESSAGE_UPDATE"
)

func ParseCmd(body []byte) string {

	reg, _ := regexp2.Compile(`(?<={"cmd":")[A-Z_0-9]+`, 0)

	cmd, err := reg.FindStringMatch(string(body))
	if err != nil {
		fmt.Println("CMD错误", err)
		return ""
	}

	return cmd.String()
}

func (c *Client) Handle(pkt Packet) {
	switch pkt.Operation {
	case NOTICE:
		//var notice Notice
		//json.Unmarshal(pkt.Body, &notice)
		CMD := ParseCmd(pkt.Body)
		switch CMD {
		case string(LINK_INFO_V3_UPDATE):
			lu := new(message.LinkUp)
			lu.Parse(pkt.Body)
		case string(LIKE_INFO_V3_UPDATE):
			lu := new(message.LikeUp)
			lu.Parse(pkt.Body)
		case string(DANMU_MSG):
			dm := new(message.DanMu)
			dm.Parse(pkt.Body)
			base := models.Base{
				Rid:  string(rune(c.roomId)),
				Type: constants.ChatMsgType,
			}
			chatMsg := models.ChatMsgMessage{
				Base:     base,
				UserInfo: dm.UserInfo,
				Txt:      dm.Extra.Content,
			}
			fmt.Println(chatMsg)
			c.publish.Publish(chatMsg.GetType(), chatMsg)
		case string(INTERACT_WORD):
			iw := new(message.InteractWord)
			iw.Parse(pkt.Body)
		case string(WATCHED_CHANGE):
			iw := new(message.WatchedChange)
			iw.Parse(pkt.Body)
		case string(SEND_GIFT):
			gift := new(message.Gift)
			gift.Parse(pkt.Body)
		case string(ONLINE_RANK_COUNT):
			iw := new(message.OnlineRankCount)
			iw.Parse(pkt.Body)
		case string(ONLINE_RANK_COUNT_V2):
		case string(ONLINE_RANK_V2):
		case string(STOP_LIVE_ROOM_LIST):
		default:
			fmt.Println(CMD, string(pkt.Body))

		}
		break
	case HEART_BEAT_RESPONSE:
	case ENTER_ROOM_RESPONSE:
	default:
		fmt.Println("unknown Operation")
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
	client.publish.AddMessageHandle(MessageHandle)
	return client
}

// 对消息信息结构
func MessageHandle(p *models.Publisher, message string) {
	fmt.Println(message)
}

// JoinRoom
// https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=27848294 获取key以及登录token
func (c *Client) JoinRoom(roomId int) error {
	resp, err := http.Get(fmt.Sprintf("https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=%d", roomId))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Body, "resp")
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response.Data.Token)
	// 获取token
	c.roomId = roomId
	auth := &AuthParams{
		Roomid:   c.roomId,
		Uid:      0,
		Protover: 3,
		Type:     2,
		Platform: "web",
		Key:      response.Data.Token,
		//ClientVer: "1.16.3",
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
			log.Panicln("Connection closed by remote side!", err)
		} else {
			log.Panicln("出错了")
		}
		return err
	}
	return nil
}
