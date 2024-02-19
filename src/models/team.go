package models

import (
	"errors"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Id      uint
	Name    string    `gorm:"unique;not null"`
	Members []*Member `gorm:"many2many:member_teams"` // Teamに所属するTeamのスライス
}

func (r *Repository) CreateTeam(name string, members *[]Member) error {

	team := Team{Name: name}

	if name == "" {
		return errors.New("team name cannot be empty")
	}

	r.db.Create(&team)

	membersSlice := *members
	for _, member := range membersSlice {
		// 新しいMemberポインタを作成して、それをteam.Membersに追加する
		newMember := member // 新しいMemberオブジェクトを作成
		team.Members = append(team.Members, &newMember)
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

func (r *Repository) RandomTeam() *Team {
	var team Team
	if err := r.db.Order("RANDOM()").Limit(1).Find(&team).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &team
}
