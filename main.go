package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/musketeer-liu/Go_Mega_Project/controller"
	"github.com/musketeer-liu/Go_Mega_Project/model"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	http.ListenAndServe(":8888", nil)
}
