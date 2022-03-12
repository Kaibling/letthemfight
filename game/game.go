package game

import (
	"bufio"
	"fmt"
	"letthemfight/module/database"
	"letthemfight/repositories"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type Game struct {
}

func (g *Game) Start() {
	db := InitGame()
	//cr := repositories.NewCharacterRepo(db)

	ur := repositories.NewUserRepo(db)
	user := repositories.NewUser("Hansi")
	user.Characters = append(user.Characters, *repositories.NewCharacter("Der SChlachter"))
	ur.Save(user)

	// //cr.AddCharacter()
	// for {
	// 	//generate Enemies
	// 	//let player choose character
	// 	cr.PrintHeroes()
	// 	input := readInput()

	// 	fight.NewFight()
	// 	//FIGHT
	// 	//repeat
	// }

}

func NewGame() *Game {
	return &Game{}
}

func InitGame() *gorm.DB {
	// Migrate the schema
	db := database.NewDB()
	migrators := []repositories.DBMigrator{}
	migrators = append(migrators, repositories.NewCharacterMigrator(db))
	migrators = append(migrators, repositories.NewGroupDBMigrator(db))
	migrators = append(migrators, repositories.NewUserDBMigrator(db))
	for i := range migrators {
		err := migrators[i].Migrate()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return db
}

func readInput() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose Hero ID: ")
	text, _ := reader.ReadString('\n')
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("no valid number")
	}
	return number
}
