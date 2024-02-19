package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Repository struct {
	db *gorm.DB
}

const databaseFile = "db/sample.db"

func NewRepository() (*Repository, error) {
	db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repository := &Repository{db: db}
	return repository, nil
}

func Init() {
	db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database\n")
	}
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Team{})
}
