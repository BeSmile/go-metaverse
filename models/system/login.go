package system

import (
	"go-metaverse/global/orm"
	"go-metaverse/tools"
)

type Login struct {
	Username  string `form:"UserName" json:"username" binding:"required"`
	Password  string `form:"Password" json:"password" binding:"required"`
	Code      string `form:"Code" json:"code"`
	UUID      string `form:"UUID" json:"uuid" binding:"required"`
	LoginType int    `form:"LoginType" json:"loginType"`
}

func (u *Login) GetUser() (user SysUser, e error) {
	e = orm.DB.Table("sys_user").Where("username = ? ", u.Username).Find(&user).Error
	if e != nil {
		return
	}
	if u.LoginType == 0 {
		_, e = tools.CompareHashAndPassword(user.Password, u.Password)
		if e != nil {
			return
		}
	}
	return
}
