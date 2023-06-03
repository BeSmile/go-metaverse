package message

import (
	"encoding/json"
	"github.com/thedevsaddam/gojsonq/v2"
)

type WatchedChange struct {
	Num       int    `json:"num"`
	TextSmall string `json:"text_small"`
	TextLarge string `json:"text_large"`
}

func (wc *WatchedChange) Parse(body []byte) {
	data := gojsonq.New().FromString(string(body)).Find("data")
	jsonData, _ := json.Marshal(data)
	json.Unmarshal(jsonData, &wc)
}
