package system

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

type Password struct {
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type LoginM struct { // 登录请求
	UserName
	Password
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT" json:"userId"`
}

type SysUserBase struct {
	NickName string `gorm:"type:varchar(64)" json:"nick_name"`
	Status   int    `gorm:"type:int(1)" json:"status"`
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
