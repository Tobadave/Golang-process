package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.Error(w, "ERROR: Page Not Found", http.StatusNotFound)
			return
		}

		fmt.Println(r.Method)
		fmt.Println(r.URL.Path)
		fmt.Fprintf(w, "Hello from Go!")

	})

	/*
		http.HandleFunc → registers route + handler
		http.ListenAndServe → starts server, waits for requests
		request comes in → Go finds matching handler → executes it → writes to w → browser gets response
	*/

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		// if r.URL.Path != "/about" {
		// 	http.Error(w, "404 not found TESTED", http.StatusNotFound)
		// }
		fmt.Fprintf(w, "WELCOME TO ABOUT SECTION")
	})

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, "You are on GET method")
		} else if r.Method == "POST" {
			fmt.Fprintf(w, "You are on POST methos")
		}
	})

	http.ListenAndServe(":8080", nil)
}
