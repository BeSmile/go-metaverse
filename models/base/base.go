package base

// 基础类所有的类会继承这个基础类，

type Model struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	CreatedAt int `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdatedAt int `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeletedAt int `gorm:"column:delete_time" json:"delete_time" form:"delete_time"`
}
