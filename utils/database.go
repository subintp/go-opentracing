package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

// GetDbConnection - Database connection
func GetDbConnection() *gorm.DB {
	if connection == nil {
		dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		connection = db
	}
	return connection
}