package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/musketeer-liu/Go_Mega_Project/model"

	"log"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})
}



