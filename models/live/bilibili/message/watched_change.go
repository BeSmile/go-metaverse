package message

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"go-metaverse/tools/bytes"
)

type WatchedChange struct {
	Num       int    `json:"num"`
	TextSmall string `json:"text_small"`
	TextLarge string `json:"text_large"`
}

func (wc *WatchedChange) Parse(body []byte) {
	data := gjson.Parse(bytes.BytesToString(body)).String()
	jsonData, _ := json.Marshal(data)
	json.Unmarshal(jsonData, &wc)
}
