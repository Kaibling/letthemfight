package repositories

import "gorm.io/gorm"

type Group struct {
	Name  string
	chars []*Character
}

func (g *Group) AddCharacter(char *Character) {
	g.chars = append(g.chars, char)
}

func NewGroup(name string) Group {
	return Group{Name: name}
}

type GroupDBMigrator struct {
	db *gorm.DB
}

func NewGroupDBMigrator(db *gorm.DB) *GroupDBMigrator {
	return &GroupDBMigrator{db}
}

func (c *GroupDBMigrator) Migrate() error {
	return c.db.AutoMigrate(&Group{})
}
