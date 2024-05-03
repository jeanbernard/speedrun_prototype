package database

import (
	"developer/any/clients/sqlite"
	"sync"

	"gorm.io/gorm"
)

type SQLiteDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *SQLiteDatabase
)

func NewSQLiteDatabase() Database {
	once.Do(func() {
		db, err := sqlite.OpenDB("test.db")
		if err != nil {
			panic("failed to connect to SQLite")
		}
		dbInstance = &SQLiteDatabase{db}
	})

	return dbInstance
}

func (s SQLiteDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
