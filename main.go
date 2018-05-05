package main

import (
	"context"

	"github.com/koodinikkarit/matias/program"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	version = "0.0.1"
)

func main() {
	ctx := context.Background()
	program := program.NewProgram(ctx)
	program.Start()
}

// if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
// 	fmt.Println("config file does not exists")
// }

// db, _ := gorm.Open("sqlite3", "database.db")
// db.Debug().AutoMigrate(
// 	&models.Variation{},
// )
// defer db.Close()

// db, _ := gorm.Open("sqlite3", "Songs.db")
// defer db.Close()

// songs := []ewDatabaseModels.Song{}

// db.Debug().Table("song").Find(&songs)

// for index, s := range songs {
// 	if s.Author != "" {
// 		fmt.Println(index, s.Title, s.Author)
// 	}
// }

// fmt.Println("tulokset")

// if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
// 	fmt.Println("config file does not exists")
// }

// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(dir, os.Args[0])
// err2 := os.Remove(os.Args[0])
// if err2 != nil {
// 	fmt.Println("error", err2)
// }
// versionNumber := 1

// fmt.Println("Versionnumber is", versionNumber)

// if versionNumber == 1 {
// 	fmt.Println("Newest version is 2")
// 	fmt.Println("Updating version...")
// 	doUpdate("http://localhost:2233")
// 	cmd := exec.Command("main.exe")
// 	cmd.Stdout = os.Stdout
// 	cmd.Start()
// }
