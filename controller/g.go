package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
)

var (
	homeController	home
	templates		map[string]*template.Template
	// 设置全局变量 sessionName 和 store
	sessionName		string
	flashName		string
	// store初始化的时候最好是从配置文件中读取 这里直接 hard code 了
	store			*sessions.CookieStore
	// 更灵活的操作室将pageLimit 放到配置文件中
	pageLimit		int
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("something_very_secret"))
	sessionName = "go_mega"
	flashName = "go_flash"
	pageLimit = 5
}

func Startup() {
	homeController.registerRoutes()
}