package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	Characters []Character
}

func NewUser(name string) *User {
	return &User{Name: name}
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Save(user *User) {
	err := r.db.Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}

// func (r *UserRepo) LoadByID(id int) {
// 	char := User{ID: id}
// 	err := r.db.Find(&char).Error
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

func (r *UserRepo) PrintHeroes() {
	var heros []Character
	err := r.db.Find(&heros).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	for i := range heros {
		fmt.Printf("id: %d, Name: %s\n", heros[i].ID, heros[i].Name)
	}
}

type UserDBMigrator struct {
	db *gorm.DB
}

func NewUserDBMigrator(db *gorm.DB) *UserDBMigrator {
	return &UserDBMigrator{db}
}

func (c *UserDBMigrator) Migrate() error {
	return c.db.AutoMigrate(&User{})
}
