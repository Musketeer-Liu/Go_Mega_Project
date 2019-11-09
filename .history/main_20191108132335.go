package main

import (
	// "fmt"
	// "io"
	"net/http"
	"html/templates"
)

type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := 


		// w.Write([]byte("Hello World!"))
		// fmt.Fprintf(w, "Hello World!")
		// io.WriteString(w, "Hello World!")

	})
	http.ListenAndServe(":8888", nil)
}
