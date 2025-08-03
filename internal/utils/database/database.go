package database

import (
	"gorm.io/gorm"
)

type Database interface {
	GetDB() *gorm.DB
	Close() error
	Migrate(models ...interface{}) error
}

type DBConfig struct {
	Driver string
	DSN    string
}
