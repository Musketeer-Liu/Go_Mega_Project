package main

import (
	// "fmt"
	// "io"
	"net/http"
	"html/template"
)

type User struct {
	Username string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "Musketeer"}
		tpl, _ := template.New("").Parse(
			`<html>
				<head>
					<title>Home Page - Bonfy</title>
				</head>
				<body>
					<h1>Hello, {{.Username}}!</h1>
				</body>
			</html>`)
		tpl.E


		// w.Write([]byte("Hello World!"))
		// fmt.Fprintf(w, "Hello World!")
		// io.WriteString(w, "Hello World!")

	})
	http.ListenAndServe(":8888", nil)
}
