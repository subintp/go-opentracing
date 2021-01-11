package utils

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

func getDbConnection() *gorm.DB {
	if connection == nil {
		dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open("postgres", dsn)
		if err != nil {
			panic("failed to connect database")
		}

		AddGormCallbacks(db)
		connection = db
	}
	return connection
}

// GetDB - Database connection
func GetDB(ctx context.Context) *gorm.DB {
	return SetSpanToGorm(ctx, getDbConnection())
}
