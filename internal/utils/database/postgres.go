package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	db *gorm.DB
}

func NewPostgresDatabase(config DBConfig) (*PostgresDatabase, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{db: db}, nil
}

func (p *PostgresDatabase) GetDB() *gorm.DB {
	return p.db
}

func (p *PostgresDatabase) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (p *PostgresDatabase) Migrate(models ...interface{}) error {
	return p.db.AutoMigrate(models...)
}
