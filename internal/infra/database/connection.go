package database

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

type IGormAdapter interface {
	AutoMigrate(migrations ...interface{}) error
	Create(model interface{}) *gorm.DB
	FindAll(model interface{}) *gorm.DB
	FindByID(model interface{}, id int) *gorm.DB
	First(model interface{}) *gorm.DB
	Close() *gorm.DB
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

func (a *GormAdapter) FindAll(model interface{}) *gorm.DB {
	return a.db.Find(model)
}

func (a *GormAdapter) FindByID(model interface{}, id int) *gorm.DB {
	return a.db.First(&model, id)
}

func (a *GormAdapter) First(model interface{}) *gorm.DB {
	return a.db.First(model)
}

func (a *GormAdapter) Create(model interface{}) *gorm.DB {
	return a.db.Create(model)
}

func (a *GormAdapter) Close() *gorm.DB {
	return nil
}
