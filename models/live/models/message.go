package models

import (
	"go-metaverse/models/live/constants"
	"go-metaverse/models/live/utils"
)

type Base struct {
	// 房间 id
	Rid string `json:"rid"`
	// 弹幕组 id
	Gid  string            `json:"gid"`
	Type constants.MsgType `json:"type"`
	Sahf string            `json:"sahf"`
}

type UserInfo struct {
	// 用户头像 https://apic.douyucdn.cn/upload/avatar_v3@S202303@Sac86debfe0be426cb91b9dcb7900d242.jpg  @需要替换成/
	Ic string `json:"ic"`
	// 发送者昵称
	Nn string `json:"nn"`
	// 用户等级
	Level string `json:"level"`
	// 贵族等级
	Nl  string `json:"nl"`
	UId string `json:"uid"`
}

type NobleNumInfoMessage struct {
	//Base
	Sum string `json:"sum"`
	Vn  string `json:"vn"`
}

//func (u NobleNumInfoMessage) GetType () MsgType {
//	return u.Base.Type
//}

// 赠送礼物 dgb
type DgbMessage struct {
	Base
	UserInfo
	// 礼物id
	Gfid string `json:"gfid"`
	// 大礼物标识:默认值为 0(表示是小礼物)
	Bg string `json:"bg"`
}

func (d DgbMessage) GetType() constants.MsgType {
	return d.Base.Type
}

// 聊天信息 chatMsg
type ChatMsgMessage struct {
	Base
	UserInfo
	Pdg string `json:"pdg"`
	// 弹幕唯一 ID
	Cid string `json:"cid"`
	Dms string `json:"dms"`
	// 弹幕文本内容
	Txt string `json:"txt"`
	// 用户等级
	Level string `json:"level"`
	// 礼物头衔:默认值 0(表示没有头衔)
	Gt string `json:"gt"`
	// 颜色:默认值 0(表示默认颜色弹幕)
	Col string `json:"col"`
	// 弹幕具体类型: 默认值 0(普通弹幕)
	Cmt string `json:"cmt"`
	// nc 贵族弹幕标识,0-非贵族弹幕,1-贵族弹幕,默认值 0
	Nc string `json:"nc"`
	// 客户端, 默认0 网页, 1 安卓 2 ios
	Ct string `json:"ct"`
}

func (c ChatMsgMessage) GetType() constants.MsgType {
	return c.Base.Type
}

func (c *ChatMsgMessage) SerializeData(response utils.Response) {
	//msgType := GetMsgTypeReflect(response["type"])

	//c.Base.Type = msgType
}

// 用户进入
type UenterMessage struct {
	Base
	UserInfo
	Txt string `json:"txt"`
}

func (u UenterMessage) GetType() constants.MsgType {
	return u.Base.Type
}

type List struct {
	Uid string `json:"uid"`
	Crk string `json:"crk"`
	// 排名变化，-1:下降，0:持平，1:上升
	Rs       string `json:"rs"`
	GoldCost string `json:"gold_cost"`
}

// 排名信息 ranklist
type RankListMessage struct {
	Base
	// 排行榜更新时间戳
	Ts string `json:"ts"`
	// 排行榜序列号
	Seq string `json:"seq"`
	// 弹幕分组
	Gid string `json:"gid"`

	// 总榜，包含的子结构体信息如下:
	ListAll List `json:"list_all"`

	//周榜，包含的子结构体信息如下:
	List List `json:"list"`
	// 日榜，包含的子结构体信息如下:
	ListDay List `json:"list_day"`
}

func (r RankListMessage) GetType() constants.MsgType {
	return r.Base.Type
}

// 超级弹幕 ssd
type SsdMessage struct {
	Base
	// 超级弹幕 id
	Sdid string `json:"sdid"`
	// 跳转房间 id
	Trid string `json:"trid"`
	// 超级弹幕的内容
	Content string `json:"content"`
	// 跳转 url
	Url string `json:"url"`
	// 客户端类型
	Clitp string `json:"clitp"`
	// 跳转类型
	Jmptp string `json:"jmptp"`
}

func (s SsdMessage) GetType() constants.MsgType {
	return s.Base.Type
}

// 房间内礼物广播
type SpbcMessage struct {
	Base
	// 礼物id
	Eid string `json:"eid"`
	// sn 赠送者昵称
	Sn string `json:"sn"`
	// 受赠者昵称
	Dn string `json:"dn"`
	// 礼物名称
	Gn string `json:"gn"`
	// gc 礼物数量
	Gc string `json:"gc"`
	// 是否有礼包(0-无礼包，1-有礼包)
	Gb string `json:"gb"`
	// 广播展现样式(1-火箭，2-飞机)
	Es string `json:"es"`
	//  赠送房间
	Drid string `json:"drid"`
	// 广播礼物类型
	Bgl string `json:"bgl"`
	// 栏目分类广播字段
	Cl2 string `json:"cl2"`
}

func (s SpbcMessage) GetType() constants.MsgType {
	return s.Base.Type
}

// 房间用户抢红包
type GgbbMessage struct {
	Base
}

func (g GgbbMessage) GetType() constants.MsgType {
	return g.Base.Type
}

// rankup 房间 top10 排行榜变换
type RankUpMessage struct {
	Base
}

func (r RankUpMessage) GetType() constants.MsgType {
	return r.Base.Type
}

// 粉丝排行榜消息
type FrankMessage struct {
	Base
	// 粉丝总人数
	Fc string `json:"fc"`
	// 徽章昵称
	Bnn string `json:"bnn"`
	// 榜单版本号
	Ver  string `json:"ver"`
	List struct {
		UserInfo
		//fim 粉丝亲密度
		Fim string `json:"fim"`
		//ic 用户头像
		Ic string `json:"ic"`
		//rg: 用户房间权限组
		Rg string `json:"rg"`
		//pg: 用户平台权限组
		Pg string `json:"pg"`
		//bl 徽章等级
		Bl string `json:"fblc"`
		//hd 扩展功能字段
		Hd string `json:"hd"`
		//ri 扩展字段，一般不使用
		Ri string `json:"ri"`
		//lev: 用户等级
		Lev string `json:"lev"`
		//sahf 扩展字段，一般不使用，可忽略
		Sahf string `json:"sahf"`
	} `json:"list"`
}

func (f FrankMessage) GetType() constants.MsgType {
	return f.Base.Type
}

type Message interface {
	GetType() constants.MsgType
	//SerializeData(val interface{})
}

//func (m Message) GetType () MsgType {
//	return m.Type
//}
