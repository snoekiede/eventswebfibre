package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func InitializeDatabaseConnection() (*gorm.DB, error) {
	dsn := ConstructDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in InitializeDatabaseConnection")
		}
	}()
	connection, err := db.DB()
	if err == nil {
		connection.SetMaxOpenConns(10)
		connection.SetMaxIdleConns(5)
		connection.SetConnMaxLifetime(time.Hour)
	} else {
		return nil, err
	}
	return db, err
}

func getEnvironmentVariableWithDefault(key string, defaultValue string) string {
	currentValue := os.Getenv(key)
	if currentValue == "" {
		return defaultValue
	} else {
		return currentValue
	}
}

func ConstructDsn() string {
	host := getEnvironmentVariableWithDefault("host", "localhost")
	user := getEnvironmentVariableWithDefault("user", "postgres")
	password := getEnvironmentVariableWithDefault("password", "Piloten2030")
	dbname := getEnvironmentVariableWithDefault("dbname", "webeventsfiber")
	port := getEnvironmentVariableWithDefault("port", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	return dsn
}
