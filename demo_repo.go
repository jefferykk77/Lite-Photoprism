package main

import (
	"fmt"
	"log"
	"photoprism-lite/repository"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func main() {
	schema := repository.InitSchema()
	// 1. 直连数据库（配好钥匙）
	db, err := sqlx.Connect("sqlite", "sqlite.db")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	// 2. 拔高了！把钥匙交给库管员
	RM := repository.NewRepoManager(db)

	// 3. 干净的调用（不再出现任何 SQL 字符串）
	if err := RM.AddPhoto("/image/.png", "a photo"); err != nil {
		log.Fatalln("Add photo Error:", err)
	}
	photo, err := RM.GetByID(1)
	if err != nil {
		log.Fatalln("Geting photo Error:", err)
	}
	fmt.Println(photo)
}
