package main

import (
	// "fmt"
	// "io"
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

type Post struct {
	User User
	Body string
}

type IndexViewModel struct {
	Title string
	User  User
	Posts []Post
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{Username: "Musketeer"}
		u2 := User{User}
		v := IndexViewModel{Title: "Homepage", User: user}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &v)

		// w.Write([]byte("Hello World!"))
		// fmt.Fprintf(w, "Hello World!")
		// io.WriteString(w, "Hello World!")

	})
	http.ListenAndServe(":8888", nil)
}
