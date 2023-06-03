package message

import (
	"encoding/json"
)

type (
	OnlineRankCount struct {
		Count int `json:"count"`
	}

	OnlineRankCountRes struct {
		Data OnlineRankCount `json:"data"`
		CMD  string          `json:"cmd"`
	}
)

func (orc *OnlineRankCount) Parse(body []byte) {
	var orcRes OnlineRankCountRes
	json.Unmarshal(body, &orcRes)
	*orc = orcRes.Data
}
