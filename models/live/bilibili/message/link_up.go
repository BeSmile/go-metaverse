package message

import "encoding/json"

type (
	LinkUp struct {
		ShowArea  int    `json:"show_area"`
		LikeIcon  string `json:"like_icon"`
		MsgType   int8   `json:"msg_type"`
		Uid       int32  `json:"uid"`
		LikeText  string `json:"like_text"`
		UName     string `json:"u_name"`
		EColor    string `json:"e_color"`
		FansMedal struct {
			TargetId         int    `json:"target_id"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int8   `json:"medal_color"`
			MedalColorStart  int32  `json:"medal_color_start"`
			MedalColorEnd    int32  `json:"medal_color_end"`
			MedalColorBorder int32  `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
			Special          string `json:"special"`
			IconId           int    `json:"icon_id"`
			AnchorRoomId     string `json:"anchor_roomid"`
			Score            int    `json:"score"`
		}
	}

	LinUpRes struct {
		Data LinkUp `json:"data"`
		CMD  string `json:"cmd"`
	}
)

func (lu *LinkUp) Parse(body []byte) {
	var luRes LinUpRes
	json.Unmarshal(body, &luRes)
	*lu = luRes.Data
}
