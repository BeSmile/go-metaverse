package database

import "gorm.io/gorm"

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
	GetConnect() string
}
