package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Init connects to db server
func Init(connection string) (*gorm.DB, error) {
	gormDB, err := gorm.Open("mysql", connection)
	if err != nil {
		return db, err
	}
	// test connection
	err = gormDB.DB().Ping()
	if err != nil {
		return db, err
	}
	db = gormDB
	return db, nil
}

// GetDB return active db connection
func GetDB() *gorm.DB {
	return db
}
