package building

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-metaverse/global/orm"
	"go-metaverse/helpers/app"
	"go-metaverse/models/base"
)

type Animal struct {
	Name   string     `gorm:"type:varchar(24)" json:"name"`
	Island IslandType `gorm:"type:integer" json:"is_land"`
	Output string     `gorm:"type:varchar(24)" json:"output"`
	base.Model
}

type IslandType int

const (
	LIsland IslandType = 1 // 隆隆岛
	MIsland IslandType = 2 // 茂茂岛
	NIsland IslandType = 3 // 亮亮岛
	SIsland IslandType = 4 // 湿湿岛
	CIsland IslandType = 5 // 潺潺岛
)

// Magic 魔物
type Magic struct {
	Name   string     `gorm:"type:varchar(24)" json:"name"`
	Island IslandType `gorm:"type:integer" json:"is_land"`
	Effect string     `gorm:"type:varchar(24)" json:"effect"`
}

func (Magic) TableName() string {
	return "building_magic"
}

func (magic *Magic) SaveMagic(c *gin.Context) {
	_ = c.BindJSON(&magic)
	orm.DB.Create(&magic)
	fmt.Printf("%v", magic)
	app.OK(c, magic, "success")

}
