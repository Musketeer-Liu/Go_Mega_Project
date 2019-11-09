package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{Username: "Musketeer"}
		u2 := User{Username: "Paladin"}

		posts := []Post{
			Post{User: u1, Body: "Beautiful day in Portland!"},
			Post{User: u2, Body: "The Avengers movie was so cool!"},
		}

		v := IndexViewModel{Title: "Homepage", User: u1, Posts: posts}
		templates := PopulateTemplates()
		templates["index.html"].Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
