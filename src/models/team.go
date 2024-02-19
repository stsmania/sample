package models

import (
	"errors"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Id      uint
	Name    string    `gorm:"unique;not null"`
	Persons []*Person `gorm:"many2many:person_teams"` // Teamに所属するTeamのスライス
}

func (r *Repository) CreateTeam(name string, people *[]Person) error {

	team := Team{Name: name}

	if name == "" {
		return errors.New("team name cannot be empty")
	}

	r.db.Create(&team)

	peopleSlice := *people
	for _, person := range peopleSlice {
		// 新しいPersonポインタを作成して、それをteam.Personsに追加する
		newPerson := person // 新しいPersonオブジェクトを作成
		team.Persons = append(team.Persons, &newPerson)
	}

	// チームを保存
	if err := r.db.Save(&team).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindTeam(id int) *Team {
	var team Team
	if err := r.db.First(&team, id).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &team
}

func (r *Repository) DeleteTeam(id int) {
	r.db.Delete(&Team{}, id)
}

func (r *Repository) AllTeam() *[]Team {
	var teams []Team
	if err := r.db.Find(&teams).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &teams
}

func (r *Repository) SampleTeam() *Team {
	var team Team
	if err := r.db.Order("RANDOM()").Limit(1).Find(&team).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &team
}
