package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SetupDatabaseSQLConnection() *gorm.DB {
	dsn := "sqlserver://sa:123@localhost:1433?database=GolangDB"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", err))
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", err))
	}
	dbSQL.Close()
}
