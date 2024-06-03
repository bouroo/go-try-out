package infrastructure

import (
	"log/slog"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite(dbName string) (dbConn *gorm.DB, err error) {
	dbConn, err = gorm.Open(sqlite.Open(dbName + "?cache=shared"))
	if err != nil {
		slog.Error("NewSqlite", "gorm.Open", err)
		return
	}
	// connection pool
	if sqlDB, err := dbConn.DB(); err == nil {
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
	}
	return
}
