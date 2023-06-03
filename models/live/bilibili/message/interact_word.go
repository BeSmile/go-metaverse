package message

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
)

type (
	InteractWord struct {
		Contribution struct {
			Grade int `json:"grade"`
		} `json:"contribution"`
		CoreUserType int `json:"core_user_type"`
		DmsCore      int `json:"dms_core"`
		FansMedal    struct {
			AnchorRoomId     int    `json:"anchor_room_id"`
			GuardLevel       int    `json:"guard_level"`
			IconId           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       int    `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Score            int    `json:"score"`
			Special          string `json:"special"`
			TargetId         int    `json:"target_id"`
		} `json:"fans_medal"`
		Identities    []int  `json:"identities"`
		IsSpread      int    `json:"is_spread"`
		MsgType       int    `json:"msg_type"`
		PrivilegeType int    `json:"privilege_type"`
		RoomId        int    `json:"roomid"`
		Score         int    `json:"score"`
		SpreadDesc    int    `json:"spread_desc"`
		SpreadInfo    string `json:"spread_info"`
		TailIcon      int    `json:"tail_icon"`
		TimeStamp     int    `json:"time_stamp"`
		TriggerTime   int16  `json:"trigger_time"`
		Uid           int    `json:"uid"`
		UName         string `json:"u_name"`
		UNameColor    string `json:"u_name_color"`
	}
)

func (iw *InteractWord) Parse(body []byte) {
	data := gojsonq.New().FromString(string(body)).Find("data")
	dataJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println("interface 转 json 错误")
		return
	}
	json.Unmarshal(dataJson, iw)
}
