package controller

import (
	"github.com/musketeer-liu/Go_Mega_Project/vm"
	"net/http"
)

type home struct {}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "index.html"
	vop := vm.IndexViewModelOp{}
	v := vop.getVM()
	templates[tpName].Execute(w, &v)
}

func check(username, pasword string) bool {
	if username == "musketeer-liu" && paswword == "c5d19806cc93103487e7b3c7b2e2342c48e931c9" {
		return true
	}
	return false
}

func loginHandler(w http.responseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()

	if r.method == http.MethodGet {
		templates[tpName].Execute(w, &w)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 6 {
			v.AddError("password must longer than 6")
		}

		if !check(username, password) {
			v.AddError("username password not correct, please input again")
		}

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
