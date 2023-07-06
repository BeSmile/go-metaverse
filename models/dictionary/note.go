package dictionary

import (
	"go-metaverse/global/orm"
	"go-metaverse/models/base"
)

type Note struct {
	UserId   uint `gorm:"type:uint" json:"userId"`
	English  string  `json:"english" gorm:"type:varchar(255)"`
	Chinese  string  `json:"chinese" gorm:"type:varchar(255)"`
	base.Model
}

func (Note) TableName() string {
	return "gm_note"
}

func (note Note) Insert()  error {
	if err := orm.DB.Table(note.TableName()).Create(&note); err != nil {
		return err.Error
	}
	return nil
}