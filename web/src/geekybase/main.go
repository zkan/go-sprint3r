package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	go http.ListenAndServe(":9090", nil)

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello My Handler: 9090</h1>")
	})
	// go routine (concurrent)
	go http.ListenAndServe(":9091", handler)

	handler2 := http.NewServeMux()
	handler2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello %s: 9091</h1>", r.URL.Query().Get("name"))
	})
	// do not put go routine here since we need to keep main func running
	http.ListenAndServe(":9092", handler2)
}
