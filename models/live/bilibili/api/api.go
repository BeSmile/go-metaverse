package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/google/go-querystring/query"
	"go-metaverse/tools/bytes"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HostList struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	WssPort int    `json:"wss_port"`
	WsPort  int    `json:"ws_port"`
}

type RoomGroup struct {
	Group            string     `json:"group"`
	BusinessID       int        `json:"business_id"`
	Token            string     `json:"token"`
	RefreshRowFactor float32    `json:"refresh_row_factor"`
	RefreshRate      int        `json:"refresh_rate"`
	MaxDelay         int        `json:"max_delay"`
	HostList         []HostList `json:"host_list"`
}

type WebInterface struct {
	IsLoading bool `json:"isLogin"`
	WbiImg    struct {
		ImgUrl string `json:"img_url"`
		SubUrl string `json:"sub_url"`
	} `json:"wbi_img"`
}

type UserInfo struct {
	Mid  int32  `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TTL     int         `json:"ttl"`
	Data    interface{} `json:"data"`
}

func GetToken(roomId int32) Response {
	resp, err := http.Get(fmt.Sprintf("https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=%d", roomId))
	if err != nil {
		log.Fatalln(err)
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}

type UserParams struct {
	Platform string `url:"platform" json:"platform"`
	Mid      int32  `url:"mid" json:"mid"`
	//Index int `url:"index"`
	//Order        string   `url:"string"`
	//OrderAvoided bool   `url:"order_avoided"`
	//PN int   `url:"pn"`
	//PS int   `url:"ps"`
	WebLocation int32  `url:"web_location" json:"web_location"`
	Token       string `url:"token" json:"token"`
}

type EnUserParams struct {
	UserParams
	Wts  string `url:"wts"`
	WRid string `url:"w_rid"`
}

// GetUserInfo https://api.bilibili.com/x/space/wbi/acc/info?mid=3026908&token=&platform=web&web_location=1550101&w_rid=16760fadf3d872771bd71a4d92700d31&wts=1685783841
func GetUserInfo(eup EnUserParams) Response {
	client := &http.Client{}

	params, _ := query.Values(eup)

	url := "https://api.bilibili.com/x/space/wbi/acc/info?" + params.Encode()
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}

// GetWebInterface 获取基本加密信息
func GetWebInterface() Response {
	resp, err := http.Get("https://api.bilibili.com/x/web-interface/nav")
	if err != nil {
		log.Fatalln(err)
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}

// GetHash https://i0.hdslb.com/bfs/wbi/b7f68cd5444d413bb404a0a6ae865fd7.png
func GetHash(url string) string {
	reg, _ := regexp2.Compile("https://i0.hdslb.com/bfs/wbi/([a-z0-9]*).png", 0)
	hash, err := reg.FindStringMatch(url)
	if err != nil {
		fmt.Println("解析hash失败")
		return ""
	}
	return hash.Groups()[1].Captures[0].String()
}

// Encrypt 获取加密参数w-rid
/**
function(t, e) {
  e || (e = {});
  var n, r, o = function(t) {
    if (t.useAssignKey)
      return {
        imgKey: t.wbiImgKey,
        subKey: t.wbiSubKey
      };
    var e = l("wbi_img_url")
      , n = l("wbi_sub_url")
      , r = e ? f(e) : t.wbiImgKey
      , o = n ? f(n) : t.wbiSubKey;
    return {
      imgKey: r,
      subKey: o
    }
  }(e), i = o.imgKey, a = o.subKey;
  if (i && a) {
    for (var c = (n = i + a, r = [],
      [46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49, 33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40, 61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11, 36, 20, 34, 44, 52].forEach((function(t) {
          n.charAt(t) && r.push(n.charAt(t))
        }
      )),
      r.join("").slice(0, 32)), s = Math.round(Date.now() / 1e3), p = Object.assign({}, t, {
      wts: s
    }), d = Object.keys(p).sort(), h = [], v = /[!'\(\)*]/g, m = 0; m < d.length; m++) {
      var y = d[m]
        , g = p[y];
      g && "string" == typeof g && (g = g.replace(v, "")),
      null != g && h.push("".concat(encodeURIComponent(y), "=").concat(encodeURIComponent(g)))
    }
    var b = h.join("&");
    return {
      w_rid: u(b + c),
      wts: s.toString()
    }
  }
  return null
}
*/
func (up UserParams) GetWrid(wi WebInterface) EnUserParams {
	charArr := []int{46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49, 33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40, 61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11, 36, 20, 34, 44, 52}
	str := fmt.Sprintf("%s%s", GetHash(wi.WbiImg.ImgUrl), GetHash(wi.WbiImg.SubUrl))
	codeArr := make([]string, 0)
	for _, index := range charArr {
		codeArr = append(codeArr, fmt.Sprintf("%c", str[index]))
	}

	codeArr = codeArr[0:32]
	now := time.Now().Unix()
	timeSta := math.Round(float64(now))
	//v, _ := query.Values(up)
	enParams := fmt.Sprintf("mid=%d&platform=%s&token=&web_location=%d&wts=%s", up.Mid, up.Platform, up.WebLocation, strconv.Itoa(int(timeSta))) + strings.Join(codeArr, "")
	// 490f246bb247a6df61c2d404f5344831
	md5New := md5.New()
	md5New.Write(bytes.StringToBytes(enParams))

	md5String := hex.EncodeToString(md5New.Sum(nil))

	ep := EnUserParams{
		UserParams: up,
		WRid:       md5String,
		Wts:        strconv.Itoa(int(timeSta)),
	}
	return ep
}
