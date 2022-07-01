package template

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-metaverse/helpers/app"
	"io/ioutil"
	"os"
)

func GetTemplateById(c *gin.Context) {
	id := c.Params.ByName("id")
	jsonFile, err := os.Open("json/" + string(id) + ".json")

	if err != nil {
		fmt.Println(err)
		app.Error(c, 404, err, "404 Not Found")
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)
		app.OK(c, result, "")
	}
	defer jsonFile.Close()
}

func get643() string {
	return ""
}
