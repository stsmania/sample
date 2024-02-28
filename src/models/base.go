package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

const databaseFile = "db/sample.db"

func NewMember() (*MemberRepository, error) {
	db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repository := &MemberRepository{db: db}
	return repository, nil
}

func NewTeam() (*TeamRepository, error) {
	db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repository := &TeamRepository{db: db}
	return repository, nil
}

func Init() {
	db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database\n")
	}
	db.AutoMigrate(&Member{})
	db.AutoMigrate(&Team{})
}
