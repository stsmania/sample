package models

import (
	"errors"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Id    uint
	Name  string  `gorm:"unique;not null"`
	Teams []*Team `gorm:"many2many:person_teams"` // Personが所属するTeamのスライス
}

func (r *Repository) CreatePerson(name string) error {
	if name == "" {
		return errors.New("team name cannot be empty")
	}

	return r.db.Create(&Person{Name: name}).Error
}

func (r *Repository) FindPerson(id int) *Person {
	var person Person
	if err := r.db.First(&person, id).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &person
}

func (r *Repository) FindPersonByIds(ids []int) []Person {
	var people []Person
	if err := r.db.Where("Id IN ?", ids).Find(&people).Error; err != nil {
		panic("failed to connect database\n")
	}
	return people
}

func (r *Repository) DeletePerson(id int) {
	r.db.Delete(&Person{}, id)
}

func (r *Repository) AllPerson() *[]Person {
	var people []Person
	if err := r.db.Find(&people).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &people
}

func (r *Repository) SamplePerson() *Person {
	var person Person
	if err := r.db.Order("RANDOM()").Limit(1).Find(&person).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &person
}

func (r *Repository) TeamPeople(id int) []*Person {
	var team Team
	if err := r.db.Preload("Persons").First(&team, id).Error; err != nil {
		panic("failed to connect database\n")
	}
	// team.Personsをpersonsにコピー
	people := make([]*Person, len(team.Persons))
	copy(people, team.Persons)
	return people
}

func (r *Repository) SampleTeamPeople(id int) *Person {
	var person Person
	if err := r.db.Raw("SELECT * FROM people p JOIN person_teams tp ON p.id = tp.person_id WHERE tp.team_id = ? ORDER BY RANDOM()", id).First(&person).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &person
}
