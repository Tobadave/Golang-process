package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "ERROR: Unspecified path", http.StatusNotFound)
			return
		}

		switch r.Method {
		case "GET":
			read, err := template.ParseFiles("templates/index1.html") //read files
			if err != nil {
				fmt.Fprintf(w, "error in file")
				http.Error(w, "Error to find file", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			read.Execute(w, read)

		case "POST":
			text := r.FormValue("text")
			banner := r.FormValue("banner")
			fmt.Fprintf(w, "text: %s, banner: %s", text, banner)

		default:
			http.Error(w, "method not allowed", http.StatusBadRequest)
		}
	})

	http.ListenAndServe(":8080", nil)
}

