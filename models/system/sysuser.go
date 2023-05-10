package system

import (
	"errors"
	"go-metaverse/global/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct { // 用户
	// 唯一id
	IdentityKey string
	UserName    string
	FirstName   string
	LastName    string
	Role        string
}

type UserName struct {
	Username string `gorm:"type:varchar(64)" json:"username"` // 为什么是type varchar(64)
}

type PassWord struct {
	Password string `gorm:"type:varchar(128)" json:"password"`
}

// LoginM 登录结构体
type LoginM struct { // 登录请求
	UserName
	PassWord
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT" json:"userId"`
}

type SysUserBase struct {
	NickName string `gorm:"type:varchar(128)" json:"nickName"`
	Email    string `gorm:"type:varchar(128)" json:"email"`
	Phone    string `gorm:"type:int(11)" json:"phone"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`
	Status   int    `gorm:"type:int(1)" json:"status"`
	CreateBy string `gorm:"type:varchar(128)" json:"createBy"`
	UpdateBy string `gorm:"type:varchar(128)" json:"updateBy"`
	Remark   string `gorm:"type:varchar(255)" json:"remark"`
	BaseModel
}

type SysUser struct {
	SysUserId
	SysUserBase
	LoginM
}

func (SysUser) TableName() string {
	return "sys_user"
}

func (user SysUser) Insert() (id int, err error) {
	if err = user.Encrypt(); err != nil {
		return
	}
	var count int
	orm.DB.Table(user.TableName()).Where("username = ? and `delete_time` IS NULL", user.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在")
		return
	}
	if err = orm.DB.Table(user.TableName()).Create(&user).Error; err != nil {
		return
	}
	id = user.UserId
	return
}

func (user *SysUser) Encrypt() (err error) {
	if user.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		user.Password = string(hash)
		return
	}
}
