package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Id      uint
	Name    string    `gorm:"unique;not null"`
	Members []*Member `gorm:"many2many:member_teams"` // Teamに所属するTeamのスライス
}

type TeamRepository struct {
	db *gorm.DB
}

func (r *TeamRepository) CreateTeam(name string, members *[]Member) error {
	team := Team{Name: name}
	if err := r.db.Create(&team).Error; err != nil {
		return err
	}

	// メンバーを関連付けて保存
	if err := r.db.Model(&team).Association("Members").Append(members); err != nil {
		return err
	}
	return nil
}

func (r *TeamRepository) FindTeam(id int) *Team {
	var team Team
	if err := r.db.First(&team, id).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &team
}

func (r *TeamRepository) AllTeam() *[]Team {
	var teams []Team
	if err := r.db.Find(&teams).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &teams
}

func (r *TeamRepository) RandomTeam() *Team {
	var team Team
	if err := r.db.Order("RANDOM()").Limit(1).Find(&team).Error; err != nil {
		panic("failed to connect database\n")
	}
	return &team
}

func (r *TeamRepository) DeleteTeam(id int) error {
	var team Team
	if err := r.db.First(&team, id).Error; err != nil {
		return err
	}

	// 中間テーブルの関連レコードを削除
	if err := r.db.Model(&team).Association("Members").Clear(); err != nil {
		return err
	}

	// Teamを削除
	if err := r.db.Unscoped().Delete(&team).Error; err != nil {
		return err
	}
	return nil
}
