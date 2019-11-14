package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	// // Local Debug Server
	// fmt.Print("Local Debug Server Starting ...")
	//http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))

	// Heroku-Postgre Server
	fmt.Print("Heroku Server Starting")
	port := os.Getenv("PORT")
	log.Println("Running on port: ", port)
	http.ListenAndServe(":"+port, context.ClearHandler(http.DefaultServeMux))
}
