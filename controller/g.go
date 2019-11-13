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
	// store初始化的时候最好是从配置文件中读取 这里直接 hard code 了
	store			*sessions.CookieStore
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	sessionName = "go_mega"
}

func Startup() {
	homeController.registerRoutes()
}