package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
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

	fmt.Print("starting ...")
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
