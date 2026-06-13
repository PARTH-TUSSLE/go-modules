package main

import (
	"fmt"
	"net/http"
)

func rootHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write([]byte("Root working, try hitting /hello?name=your_name"))
	
}

func helloHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte("Only GET is allowed"))
		return
	}
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	_, _ = w.Write([]byte("Hello "+name))
}

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}