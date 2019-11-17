package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func init() {
	projectName := "Go_Mega_Project"
	dbType := GetDBType()
	log.Println("OS DBTYPE:", dbType)

	if IsHeroku() {
		log.Println("Get Env from os.env")
	} else {
		log.Println("Init viper")
		getConfig(projectName)
	}
}

func getConfig(projectName string) {
	//name of config file (without extension)
	viper.SetConfigName("config")

	// optionally look for config in the working dir
	viper.AddConfigPath(".")
	// call multiple times to add many search paths
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))
	// path to search for the config file
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	// Find and read the config file and handle errors
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// GetMysqlConnectingString func
func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true&loc=Local", usr, pwd, host, db, charset)
}

// GetHerokuCOnnectingString struct
func GetHerokuConnectingString() string {
	return os.Getenv("DATABASE_URL")
}

// GetSMTPConfig func
func GetSMTPConfig() (server string, port int, user, pwd string) {
	if IsHeroku() {
		server = os.Getenv("F")
		port, _ = strconv.Atoi(os.Getenv("MAIL_SMTP_PORT"))
		user = os.Getenv("MAIL_USER")
		pwd = os.Getenv("MAIL_PASSWORD")
		return
	}

	server = viper.GetString("mail.smtp")
	port = viper.GetInt("mail.smtp-port")
	user = viper.GetString("mail.user")
	pwd = viper.GetString("mail.password")
	return
}

// GetServerURL func
func GetServerURL() (url string) {
	if IsHeroku() {
		url = os.Getenv("SERVER_URL")
		return
	}
	url = viper.GetString("server.url")
	return
}

// GetDBType func
func GetDBType() string {
	dbtype := os.Getenv("DBTYPE")
	return dbtype
}

// IsHeroku func
func IsHeroku() bool {
	return GetDBType() == "heroku"
}