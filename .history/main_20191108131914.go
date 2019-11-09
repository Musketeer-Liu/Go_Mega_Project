package main

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		fmt.Fprintf(w, "Hello World!")
		io.f(w, "Hello World!")

	})
	http.ListenAndServe(":8888", nil)
}
