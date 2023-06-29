package gorm

import (
	"go-metaverse/models/building"
	"go-metaverse/models/dictionary"
	"go-metaverse/models/system"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error { // gorm模型自动注入
	//db.SingularTable(true)
	return db.AutoMigrate(
		&system.SysUser{},
		&building.Animal{},
		&dictionary.Dictionary{},
	)
}
