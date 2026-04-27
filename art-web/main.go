package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		fmt.Println(r.URL.Path)
		fmt.Fprintf(w, "Hello from Go!")

		if r.URL.Path != "/about" || r.URL.Path != "/" {
			http.Error(w, "error", http.;)
		}	
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/about" {
			http.Error(w, "404 not found TESTED", http.StatusNotFound)
		}
	})

	http.ListenAndServe(":8080", nil)
}
