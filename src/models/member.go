package models

import (
	"errors"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Id    uint
	Name  string  `gorm:"unique;not null"`
	Teams []*Team `gorm:"many2many:member"` // Memberが所属するTeamのスライス
}

func (r *Repository) CreateMember(name string) error {
	if name == "" {
		return errors.New("team name cannot be empty")
	}

	return r.db.Create(&Member{Name: name}).Error
}

func (r *Repository) FindMember(id int) *Member {
	var member Member
	if err := r.db.First(&member, id).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}

func (r *Repository) FindMemberByIds(ids []int) []Member {
	var members []Member
	if err := r.db.Where("Id IN ?", ids).Find(&members).Error; err != nil {
		panic("failed to connect database\n")
	}
	return members
}

func (r *Repository) DeleteMember(id int) {
	r.db.Delete(&Member{}, id)
}

func (r *Repository) AllMember() *[]Member {
	var member []Member
	if err := r.db.Find(&member).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}

func (r *Repository) RandomMember() *Member {
	var member Member
	if err := r.db.Order("RANDOM()").Limit(1).Find(&member).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}

func (r *Repository) TeamMembers(id int) []*Member {
	var team Team
	if err := r.db.Preload("Members").First(&team, id).Error; err != nil {
		panic("failed to connect database\n")
	}
	// team.Membersをmemberにコピー
	member := make([]*Member, len(team.Members))
	copy(member, team.Members)
	return member
}

func (r *Repository) RandomTeamMember(id int) *Member {
	var member Member
	if err := r.db.Raw("SELECT * FROM members m JOIN member_teams tm ON m.id = tm.member_id WHERE tm.team_id = ? ORDER BY RANDOM()", id).First(&member).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}
