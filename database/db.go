package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitDB инициализация подключения к БД
func InitDB(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
