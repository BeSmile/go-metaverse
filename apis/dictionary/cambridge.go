package dictionary

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"go-metaverse/helpers/app"
	dictionary "go-metaverse/models/dictionary"
	bytesUtils "go-metaverse/tools/bytes"
	"log"
	"net/http"
)

type Example struct {
	Eg    string `json:"eg"`
	Trans string `json:"trans"`
}

type Block struct {
	Trans    string    `json:"trans"`
	Examples []Example `json:"examples"`
}

type Dsense struct {
	Header  string  `json:"header"`
	English string  `json:"english"`
	Blocks  []Block `json:"blocks"`
}

func GetExplain(c *gin.Context) {
	word := c.Param("word")
	dc := dictionary.Dictionary{}
	exist := dc.ExistWord(word)
	if exist {
		dc.GetWordByName(word)
		dsen := Dsense{
			English: word,
			Header:  dc.Header,
		}
		blocks := make([]Block, 0)
		json.Unmarshal(bytesUtils.StringToBytes(dc.Blocks), &blocks)
		dsen.Blocks = blocks
		app.OK(c, dsen, "success")
		return
	}

	res, err := http.Get(fmt.Sprintf("https://dictionary.cambridge.org/dictionary/english-chinese-simplified/%s", word))

	if err != nil {
		log.Fatalln("获取文档失败")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s %s", res.StatusCode, res.Status, res.Body)
		app.Error(c, 502, errors.New("dictionary.cambridge.org发生了未知错误!"), "")
		return
	}

	// Load the HTML document
	doc, err1 := goquery.NewDocumentFromReader(res.Body)
	if err1 != nil {
		log.Fatal("读取文档流失败", err1)
	}

	dsense := Dsense{
		English: word,
		Blocks:  make([]Block, 0),
	}

	doc.Find("div.pr.dsense ").Each(func(i int, selection *goquery.Selection) {
		// 磁性标题
		h3 := selection.Find("h3.dsense_h").Text()
		zh := selection.Find("div.def-body.ddef_b>span.trans.dtrans.dtrans-se.break-cj").Text()
		dsense.Header = h3
		dsense.Header = h3

		block := Block{
			// 中文翻译
			Trans:    zh,
			Examples: make([]Example, 0),
		}

		dexamps := selection.Find(".examp.dexamp")
		dexamps.Each(func(j int, selection *goquery.Selection) {
			example := Example{
				Eg:    selection.Find("span.eg.deg").Text(),
				Trans: selection.Find("span.trans.dtrans.dtrans-se.hdb.break-cj").Text(),
			}
			block.Examples = append(block.Examples, example)
		})

		dsense.Blocks = append(dsense.Blocks, block)
	})
	dc.Header = dsense.Header
	dc.English = dsense.English
	byteBlocks, err2 := json.Marshal(dsense.Blocks)
	if err2 != nil {
		app.Error(c, 500, err2, "保存失败")
		return
	}
	dc.Type = dictionary.Cambridge
	dc.Blocks = bytesUtils.BytesToString(byteBlocks)
	dc.Insert()
	app.OK(c, dsense, "success")
}
