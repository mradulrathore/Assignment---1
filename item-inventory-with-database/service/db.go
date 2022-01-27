package service

import (
	"fmt"
	"log"

	"github.com/mradulrathore/item-inventory/config"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//go:generate mockgen -source=db.go -destination=db_mock.go -package=service
type DB interface {
	GetItems() (items []Item, err error)
}

type Repository struct {
	db *gorm.DB
}

type ListItems struct {
	Items []Item
}

func Open(cfg config.AppConfig) (*gorm.DB, func(), error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not set sql.DB params")
	}

	sqlDB.SetConnMaxIdleTime(cfg.Database.MaxConnectionIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.Database.MaxConnectionLifeTime)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConnections)

	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed to close db connections %v", err)
		}
	}

	return db, cleanup, nil
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (c Repository) GetItems() (*ListItems, error) {
	var items []Item
	if err := c.db.Find(&items).Error; err != nil {
		return nil, err
	}

	return &ListItems{Items: items}, nil
}
