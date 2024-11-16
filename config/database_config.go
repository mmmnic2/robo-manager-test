package config

import (
	"device-manager/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dbUser := "postgres"
	dbPass := "cmycxcdbc"
	dbHost := "localhost"
	dbName := "device-manager"
	dbPort := 5432
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	if err := db.AutoMigrate(&model.Robot{}); err != nil {
		panic("Error during migration: %v\n")
	}
	return db
}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Unable to get DB object from GORM:")
	}
	err = sqlDB.Close()
	if err != nil {
		panic("Unable to close the database connection:")
	}
}
