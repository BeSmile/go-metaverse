package message

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"go-metaverse/models/live/models"
	"go-metaverse/tools/bytes"
)

type (
	DanMu struct {
		Content   string
		Extra     *Extra
		Emoticon  *Emoticon
		Type      int
		Timestamp int64
		Raw       string
		UserInfo  models.UserInfo
	}
	Emoticon struct {
		BulgeDisplay   int    `json:"bulge_display"`
		EmoticonUnique string `json:"emoticon_unique"`
		Height         int    `json:"height"`
		InPlayerArea   int    `json:"in_player_area"`
		IsDynamic      int    `json:"is_dynamic"`
		Url            string `json:"url"`
		Width          int    `json:"width"`
	}
	CommonNoticeDanmaku struct {
		ContentSegments []struct {
			FontColor string `json:"font_color"`
			Text      string `json:"text"`
			Type      int    `json:"type"`
		} `json:"content_segments"`
		DmsCore   int   `json:"dmscore"`
		Terminals []int `json:"terminals"`
	}
	Extra struct {
		SendFromMe     bool        `json:"send_from_me"`
		Mode           int         `json:"mode"`
		Color          int32       `json:"color"`
		DmType         int         `json:"dm_type"`
		FontSize       int         `json:"font_size"`
		PlayerMode     int         `json:"player_mode"`
		ShowPlayerType int         `json:"show_player_type"`
		Content        string      `json:"content"`
		UserHash       string      `json:"user_hash"`
		EmoticonUnique string      `json:"emoticon_unique"`
		Direction      int         `json:"direction"`
		PkDirection    int         `json:"pk_direction"`
		JumpToUrl      string      `json:"jump_to_url"`
		SpaceType      string      `json:"space_type"`
		SpaceUrl       string      `json:"space_url"`
		Animation      interface{} `json:"animation"`
		IsAudited      bool        `json:"is_audited"`
		IdStr          string      `json:"id_str"`
	}
)

func (dm *DanMu) Parse(body []byte) {
	danMuMsg := bytes.BytesToString(body)
	var extra Extra
	var emo Emoticon
	var userInfo models.UserInfo
	info := gjson.Parse(danMuMsg).Get("info")

	json.Unmarshal(bytes.StringToBytes(info.Get("0.15.extra").String()), &extra)
	json.Unmarshal(bytes.StringToBytes(info.Get("0.13").String()), &emo)
	userInfo.UId = info.Get("2.0").String()
	userInfo.Nn = info.Get("2.1").String()
	userInfo.Level = info.Get("4.1").String()
	dm.Extra = &extra
	dm.UserInfo = userInfo
	dm.Emoticon = &emo
}
