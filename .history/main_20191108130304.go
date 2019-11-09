package main


import "net/http"


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)) {
		w.Writer([]byte("Hello World!"))
	})
	http.ListenAndServe(":8888", nil)
}
