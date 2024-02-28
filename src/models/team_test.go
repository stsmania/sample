package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("team_test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Team{})

	return db
}

func TestTeamCreation(t *testing.T) {
	db := SetupDatabase()

	// cleanup after test
	defer db.Migrator().DropTable(&Team{})

	// test cases
	tests := map[string]struct {
		input Team
		want  error
	}{
		"SuccessfulCreation": {Team{Name: "GoDevelopers", Members: []*Member{}}, nil},
		"DuplicateName":      {Team{Name: "GoDevelopers", Members: []*Member{}}, gorm.ErrRecordNotFound}, // expected error is a placeholder
		"NullName":           {Team{Name: "", Members: []*Member{}}, gorm.ErrRecordNotFound},               // expected error is a placeholder
		// add more tests cases to cover other scenarios
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := db.Create(&tc.input)
			if result.Error != tc.want {
				t.Fatalf("got %v, want %v", result.Error, tc.want)
			}
		})
	}
}