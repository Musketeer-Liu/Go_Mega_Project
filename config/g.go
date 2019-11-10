package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	projectName := "go_mega"
	getConfig(projectName)
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

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=true", usr, pwd, host, db, charset)
}
