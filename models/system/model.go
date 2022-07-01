package system

import (
	"time"
)

//system模块的基础模块
type BaseModel struct {
	CreatedAt time.Time  `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdatedAt time.Time  `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeletedAt *time.Time `gorm:"column:delete_time" sql:"index" json:"-"` // 这里为什么会是指针
}
