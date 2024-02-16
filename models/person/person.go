package person

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type Person struct {
    gorm.Model
    Id   uint
    Name string
}

const databaseFile = "db/sample.db"

func Db_init() {
        db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
        if err != nil {
                panic("failed to connect database\n")
        }

        db.AutoMigrate(&Person{})
}
func Create(name string) {
        db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
        if err != nil {
                panic("failed to connect database\n")
        }
        db.Create(&Person{Name: name})
}

func Db_find(id int) Person {
        db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
        if err != nil {
                panic("failed to connect database\n")
        }
        var person Person
   	    db.First(&person, id)
        return person
}

func Db_delete(id int) {
        db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
        if err != nil {
                panic("failed to connect database\n")
        }
   	    db.Delete(&Person{}, id)
}

func Get_all() []Person {
        db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
        if err != nil {
                panic("failed to connect database\n")
        }
        var people []Person
        db.Find(&people)
        return people
}

func Sample() Person {
        db, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
        if err != nil {
            panic("failed to connect database\n")
        }


        var person Person
        db.Order("RANDOM()").Limit(1).Find(&person)
        return person
}
