package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name string
}

func NewCharacter(name string) *Character {
	return &Character{Name: name}
}

type CharRepo struct {
	db *gorm.DB
}

func NewCharacterRepo(db *gorm.DB) *CharRepo {
	return &CharRepo{db}
}

func (c *CharRepo) AddCharacter(char *Character) {
	err := c.db.Create(&char).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}

// func (c *CharRepo) LoadCharacterByID(id int) {
// 	char := Character{ID: uint(id)}
// 	err := c.db.Find(&char).Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

type CharacterDBMigrator struct {
	db *gorm.DB
}

func NewCharacterMigrator(db *gorm.DB) *CharacterDBMigrator {
	return &CharacterDBMigrator{db}
}

func (c *CharacterDBMigrator) Migrate() error {
	return c.db.AutoMigrate(&User{})
}
