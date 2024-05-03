package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB(dbname string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("failed to connect to SQLite")
	}
	return db, nil
}
