package dictionary

import (
	"go-metaverse/global/orm"
	"go-metaverse/models/base"
)

type DType int

const (
	Cambridge DType = 1
)

type Dictionary struct {
	English string `gorm:"type:varchar(255)" json:"english"`
	Type    DType  `gorm:"type:int(1)" json:"type"`
	Header  string `gorm:"type:varchar(255)" json:"header"`
	Blocks  string `gorm:"type:json" json:"blocks"`
	base.Model
}

func (Dictionary) TableName() string {
	return "gm_dictionary"
}

func (d Dictionary) Insert() error {
	if err := orm.DB.Table(d.TableName()).Create(&d); err != nil {
		return err.Error
	}
	return nil
}

func (d Dictionary) ExistWord(name string) bool {
	var count int64
	orm.DB.Model(&d).Where("english = ?", name).Count(&count)
	return count != 0
}
func (d *Dictionary) GetWordByName(name string) {
	orm.DB.Find(&d, "english = ?", name)
}
