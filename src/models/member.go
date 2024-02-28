package models

import (
	"errors"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Id    uint
	Name  string  `gorm:"unique;not null"`
	Teams []*Team `gorm:"many2many:member_teams"` // Memberが所属するTeamのスライス
}

type MemberRepository struct {
	db *gorm.DB
}

func (r *MemberRepository) CreateMember(name string) error {
	if name == "" {
		return errors.New("member name cannot be empty")
	}

	return r.db.Create(&Member{Name: name}).Error
}

func (r *MemberRepository) FindMember(id int) *Member {
	var member Member
	if err := r.db.First(&member, id).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}

func (r *MemberRepository) FindMemberByIds(ids []int) []Member {
	var members []Member
	if err := r.db.Where("Id IN ?", ids).Find(&members).Error; err != nil {
		panic("failed to connect database\n")
	}
	return members
}

func (r *MemberRepository) DeleteMember(id int) {
	r.db.Unscoped().Delete(&Member{}, id)
}

func (r *MemberRepository) AllMember() *[]Member {
	var member []Member
	if err := r.db.Find(&member).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}

func (r *MemberRepository) RandomMember() *Member {
	var member Member
	if err := r.db.Order("RANDOM()").Limit(1).Find(&member).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}

func (r *MemberRepository) TeamMembers(id int) ([]*Member, error) {
	var team Team
	if err := r.db.Preload("Members").First(&team, id).Error; err != nil {
		return nil, err
	}

	return team.Members, nil
}

func (r *MemberRepository) RandomTeamMember(id int) *Member {
	var member Member
	if err := r.db.Raw("SELECT * FROM members m JOIN member_teams tm ON m.id = tm.member_id WHERE tm.team_id = ? ORDER BY RANDOM()", id).First(&member).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &member
}
