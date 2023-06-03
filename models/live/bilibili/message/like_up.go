package message

import (
	"encoding/json"
)

type (
	LikeUp struct {
		ClickCount int16 `json:"click_count"`
	}

	LikeUpRes struct {
		Data LikeUp `json:"data"`
		CMD  string `json:"cmd"`
	}
)

func (orc *LikeUp) Parse(body []byte) {
	var orcRes LikeUpRes
	json.Unmarshal(body, &orcRes)
	*orc = orcRes.Data
}
