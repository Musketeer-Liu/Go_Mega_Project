package main

import (
	"fmt"
	"io"
	"net/http"
)



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		fmt.Fprintf(w, "Hello World!")
		io.WriteString(w, "Hello World!")

	})
	http.ListenAndServe(":8888", nil)
}
