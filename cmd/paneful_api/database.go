package main

import (
	"log"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type DBClient interface {
	Ready() bool
	RunMigration() error
}

type Client struct {
	db *gorm.DB
}

func (c Client) Ready() bool {
	var ready string
	result := c.db.Raw("SELECT 1 as ready").Scan(&ready)
	if result.Error != nil {
		return false
	}
	return ready == "1"
}

func (c Client) RunMigration() error {
	if !c.Ready() {
		log.Fatal("Database is not ready")
	}
	err := c.db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}

func NewDBClient() (Client, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	databasePort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Invalid database port")
	}
}

// refraction
// refractor
// refract
// go-reflect
