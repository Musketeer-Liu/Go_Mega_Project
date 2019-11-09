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

type 

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "Musketeer"}
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &user)

		// w.Write([]byte("Hello World!"))
		// fmt.Fprintf(w, "Hello World!")
		// io.WriteString(w, "Hello World!")

	})
	http.ListenAndServe(":8888", nil)
}
