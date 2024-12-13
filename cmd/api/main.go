package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./static/index.html")

		if err != nil {
			log.Println(err)
			w.WriteHeader(404)
			w.Write([]byte("Not Found"))
			return
		}

		tmpl.Execute(w, struct{}{})
	})

	fmt.Println("Server running http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
