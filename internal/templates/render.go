package templates

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, path string, data any) {
	tmpl, err := template.ParseFiles(
		fmt.Sprintf("./templates/%s.html", path),
	)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, data)
}
