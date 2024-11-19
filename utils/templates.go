package utils

import (
	"html/template"
	"log"
	"net/http"
)

func CreateTemplate(tmplPath string) (*template.Template, error) {
	tmpl, err := template.ParseFiles("templates/base.html", tmplPath)
	if err != nil {
		return nil, err
	}

	return tmpl, nil

}

func RenderTemplate(w http.ResponseWriter, tmplPath string, data interface{}) {
	tmpl, err := CreateTemplate(tmplPath)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

}
