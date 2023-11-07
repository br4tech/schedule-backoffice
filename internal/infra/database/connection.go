package database

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

type IGormAdapter interface {
	AutoMigrate(migrations ...interface{}) error
	Create(model any) *gorm.DB
	FindAll(model any) *gorm.DB
	FindByID(model any, id int) *gorm.DB
	First(model any) *gorm.DB
}

func NewGormAdapter(db *gorm.DB) IGormAdapter {
	return &GormAdapter{db: db}
}

func (a *GormAdapter) AutoMigrate(migrations ...interface{}) error {
	return a.db.AutoMigrate(migrations...)
}

func (a *GormAdapter) GetDB() *gorm.DB {
	return a.db
}

func (a *GormAdapter) FindAll(model any) *gorm.DB {
	return a.db.Find(model)
}

func (a *GormAdapter) FindByID(model any, id int) *gorm.DB {
	return a.db.First(&model, id)
}

func (a *GormAdapter) First(model any) *gorm.DB {
	return a.db.First(model)
}

func (a *GormAdapter) Create(model any) *gorm.DB {
	return a.Create(model)
}
