package gorm

import (
	"github.com/jinzhu/gorm"
	"go-metaverse/models/building"
	"go-metaverse/models/system"
)

func AutoMigrate(db *gorm.DB) error { // gorm模型自动注入
	db.SingularTable(true)
	return db.AutoMigrate(
		new(system.SysUser),
		new(building.Animal),
	).Error
}
