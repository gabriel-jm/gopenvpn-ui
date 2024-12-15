package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabriel-jm/gopenvpn-ui/internal/connections"
	"github.com/gabriel-jm/gopenvpn-ui/internal/database"
	"github.com/gabriel-jm/gopenvpn-ui/internal/templates"
)

func main() {
	err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB Connected!")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			address := r.FormValue("address")
			username := r.FormValue("username")
			password := r.FormValue("password")

			connections.StablishConnection(address, username, password)
		}

		templates.RenderTemplate(w, "index", nil)
	})

	http.HandleFunc("/conn", func(w http.ResponseWriter, r *http.Request) {
		templates.RenderTemplate(w, "add-connection", nil)
	})

	fmt.Println("Server running http://localhost:8000")
	err = http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
