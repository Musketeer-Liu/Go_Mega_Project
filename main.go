package main

import (
	"github.com/musketeer-liu/Go_Mega_Project/controller"
	"net/http"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
