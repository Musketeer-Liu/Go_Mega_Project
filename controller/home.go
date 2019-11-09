package controller

import (
	"net/http"
	"github.com/musketeer-liu/Go_Mega_Project/vm"
)

type home struct {}

func (h home) registerRoutes() {
	http.handleFunc("/", indexHandler)
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.getVM()
	templates["index.html"].Execute(w, &v)
}