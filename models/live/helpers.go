package live

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

const CLIENT_MSG_TYPE = 689
const RESERVED_DATA_FIELD = 0

func ToBytes(data int, byteLength int, isLittleEndian bool) []byte {
	var dataLength uint32
	switch byteLength {
	case 2:
		dataLength = uint32(uint16(data))
	case 4:
		dataLength = uint32(data)
	default:
		panic(fmt.Errorf("Unsupported byte length: %d", byteLength))
	}

	headArray := make([]byte, byteLength)
	if isLittleEndian {
		switch byteLength {
		case 2:
			binary.LittleEndian.PutUint16(headArray, uint16(dataLength))
		case 4:
			binary.LittleEndian.PutUint32(headArray, dataLength)
		}
	} else {
		switch byteLength {
		case 2:
			binary.BigEndian.PutUint16(headArray, uint16(dataLength))
		case 4:
			binary.BigEndian.PutUint32(headArray, dataLength)
		}
	}
	return headArray
}

func IntToBytes(bys int32, byteorder string) []byte {
	bytebuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytebuffer, binary.BigEndian, bys)
	if err != nil {
		return nil
	}
	BytesList := bytebuffer.Bytes()
	switch byteorder != "" {
	case byteorder == "little":
		for i := 0; i < len(BytesList)/2; i++ {
			li := len(BytesList) - i - 1
			BytesList[i], BytesList[li] = BytesList[li], BytesList[i]
		}
	}
	return BytesList
}

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
	lenByte := IntToBytes(dataLen, "little")
	sendByte := []byte{0xb1, 0x02, 0x00, 0x00}
	endByte := []byte{0x00}
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
		contentLength := BytesToInt(msgBytes[pos:pos+4], "little")
		content := bytes.NewBuffer(msgBytes[pos+12 : pos+3+contentLength]).String()
		msg = append(msg, content)
		pos = 4 + contentLength + pos
	}
	return msg
}

// BytesToInt BytesToInt:byte组转数字
func BytesToInt(bys []byte, byteOrder string) int {
	switch byteOrder != "" {
	case byteOrder == "little":
		for i := 0; i < len(bys)/2; i++ {
			li := len(bys) - i - 1
			bys[i], bys[li] = bys[li], bys[i]
		}
	}
	byteBuff := bytes.NewBuffer(bys)
	var data int32
	binary.Read(byteBuff, binary.BigEndian, &data)
	return int(data)
}

type Response map[string]interface{}

func ParseMsg(rawMsg string) Response {
	res := make(Response)
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

func GetMsgTypeReflect(val interface{}) MsgType {
	str, ok := val.(string)
	if !ok {
		fmt.Println("failed to convert to string", val)
	}
	var msgType MsgType
	switch str {
	case string(ChatMsgType):
		msgType = ChatMsgType
		break
	case string(DgbType):
		msgType = DgbType
		break
	case string(NobleNumInfoType):
		msgType = NobleNumInfoType
		break
	case string(UEnterType):
		msgType = UEnterType
		break
	case string(RankListType):
		msgType = RankListType
		break
	case string(SsdType):
		msgType = SsdType
		break
	case string(SpbcType):
		msgType = SpbcType
		break
	case string(RankUpType):
		msgType = RankUpType
		break
	case string(GgbbType):
		msgType = GgbbType
		break
	case string(FrankType):
		msgType = FrankType
		break
		// 处理其它枚举常量
	}
	return msgType
}

const Domain string = "https://apic.douyucdn.cn/upload"

func GetDomainAvatar(avatar string) string {
	url := fmt.Sprintf("%s/%s_middle.jpg", Domain, avatar)
	newUrl := strings.Replace(url, "@S", "/", -1)
	newUrl2 := strings.Replace(newUrl, "@A", "@", -1)
	return newUrl2
}
