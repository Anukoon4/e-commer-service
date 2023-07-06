package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

type (
	MySQL struct {
	}
)

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("[mysql] failed to connect database: %s", err))
	}
	if c, err := db.DB(); err != nil {
		panic(fmt.Sprintf("[mysql] connection poll error: %s", err))
	} else {
		c.SetMaxIdleConns(10)
	}
	db = db.Debug()

	DB = db
	return DB
}

func (c *MySQL) DB() *gorm.DB {
	return DB
}
