package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go-metaverse/models/live/constants"
	"strings"
)

const CLIENT_MSG_TYPE = 689
const RESERVED_DATA_FIELD = 0

// IntToBytes
func IntToBytes(bys int32, byteorder string) []byte {
	bytebuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytebuffer, binary.BigEndian, bys)
	if err != nil {
		return nil
	}
	BytesList := bytebuffer.Bytes()
	// 大端字节序中，高位字节存储在内存的低地址处，低位字节存储在内存的高地址处；而在小端字节序中，则正好相反，低位字节存储在内存的低地址处，高位字节存储在内存的高地址处
	// go默认使用的是大端,
	switch byteorder {
	case "little":
		for i := 0; i < len(BytesList)/2; i++ {
			li := len(BytesList) - i - 1
			BytesList[i], BytesList[li] = BytesList[li], BytesList[i]
		}
		break
	default:
	}
	return BytesList
}

func Int16ToBytes(bys int16, byteorder string) []byte {
	bytebuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytebuffer, binary.BigEndian, bys)
	if err != nil {
		return nil
	}
	BytesList := bytebuffer.Bytes()
	switch byteorder {
	case "little":
		for i := 0; i < len(BytesList)/2; i++ {
			li := len(BytesList) - i - 1
			BytesList[i], BytesList[li] = BytesList[li], BytesList[i]
		}
		break
	default:
	}
	return BytesList
}

// BytesToInt BytesToInt:byte组转数字
func BytesToInt(bys []byte, byteOrder string) int {
	switch byteOrder {
	case "little":
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

func GetMsgTypeReflect(val interface{}) constants.MsgType {
	str, ok := val.(string)
	if !ok {
		fmt.Println("failed to convert to string", val)
	}
	var msgType constants.MsgType
	switch str {
	case string(constants.ChatMsgType):
		msgType = constants.ChatMsgType
		break
	case string(constants.DgbType):
		msgType = constants.DgbType
		break
	case string(constants.NobleNumInfoType):
		msgType = constants.NobleNumInfoType
		break
	case string(constants.UEnterType):
		msgType = constants.UEnterType
		break
	case string(constants.RankListType):
		msgType = constants.RankListType
		break
	case string(constants.SsdType):
		msgType = constants.SsdType
		break
	case string(constants.SpbcType):
		msgType = constants.SpbcType
		break
	case string(constants.RankUpType):
		msgType = constants.RankUpType
		break
	case string(constants.GgbbType):
		msgType = constants.GgbbType
		break
	case string(constants.FrankType):
		msgType = constants.FrankType
		break
		// 处理其它枚举常量
	}
	return msgType
}

func GetDomainAvatar(avatar string) string {
	url := fmt.Sprintf("%s/%s_middle.jpg", constants.AvatarDomain, avatar)
	newUrl := strings.Replace(url, "@S", "/", -1)
	newUrl2 := strings.Replace(newUrl, "@A", "@", -1)
	return newUrl2
}
