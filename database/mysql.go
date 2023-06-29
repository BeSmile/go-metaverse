package database

import (
	"bytes"
	"fmt"
	orm "go-metaverse/global/orm"
	"go-metaverse/tools/config"
	"gorm.io/driver/mysql" //加载mysql
	"gorm.io/gorm"
	"strconv"
)

var (
	Dbtype   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
)

type Mysql struct {
}

func (e *Mysql) Setup() {
	var err error
	var database Database

	database = new(Mysql)
	orm.MysqlConn = database.GetConnect()
	orm.DB, err = database.Open(Dbtype, orm.MysqlConn)

	if err != nil {
		fmt.Println("connect error:", err)
	} else {
		fmt.Println("connect success:", Dbtype)
	}
}

func (e *Mysql) GetConnect() string {
	Dbtype = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password
	Name = config.DatabaseConfig.Name

	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=10000ms")
	return conn.String()
}

func (e *Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(conn), &gorm.Config{})
}
