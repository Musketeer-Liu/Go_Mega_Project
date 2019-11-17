package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/musketeer-liu/Go_Mega_Project/config"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	if config.IsHeroku() {
		return ConnectToDBByDBType("postgres", config.GetHerokuConnectingString())
	}
	return ConnectToDBByDBType("mysql", config.GetMysqlConnectingString())
}

// ConnectToDBByDBType func
func ConnectToDBByDBType(dbtype, connectingStr string) *gorm.DB {
	log.Println("DB Type:", dbtype, "\n Cnnect to db...")
	db, err := gorm.Open(dbtype, connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}