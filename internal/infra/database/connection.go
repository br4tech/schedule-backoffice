package database

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
  db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) *GormAdapter {
    return &GormAdapter{db: db}
}

func(a *GormAdapter) GetDB() *gorm.DB{
    return a.db
}
