package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDatabase struct {
	db *gorm.DB
}

func NewMySQLDatabase(config DBConfig) (*MySQLDatabase, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MySQLDatabase{db: db}, nil
}

func (m *MySQLDatabase) GetDB() *gorm.DB {
	return m.db
}

func (m *MySQLDatabase) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (m *MySQLDatabase) Migrate(models ...interface{}) error {
	return m.db.AutoMigrate(models...)
}
