package controller

import (
	"log"
	"net/http"

	"github.com/musketeer-liu/Go_Mega_Project/model"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		log.Println("middle:", username)
		//我们一般是不会在 controller 层直接和 model 层打交道，一般会通过 vm 层去处理，
		// 但是由于 middle 层不具有具体的view，我们这里破例直接调用 model 中的方法。
		// 也可以建立一个 middle 的vm，在里面新建一个 UpdateLastSeen 再在 controller/middle 中调用
		if username != "" {
			log.Println("Last seen:", username)
			model.UpdateLastSeen(username)
		}

		if err != nil {
			log.Println("middle get session err and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

func middleLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
