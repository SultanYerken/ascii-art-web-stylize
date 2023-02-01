package main

import (
	"ascii-art-web/cmd/ascii-art"
	"ascii-art-web/cmd/ascii-art/funcs"
	"errors"
	"html/template"
	"log"
	"net/http"
)

type News struct {
	Body string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file := "./templates/html/index.html"

	ts, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Interal Server Error", http.StatusInternalServerError)
	}
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	file := "./templates/html/index.html"

	ts, err := template.ParseFiles(file)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// writes the entered data from "name" to the variables
	// записыввает в переменные введенные данные из "name"
	text := r.FormValue("text")
	font := r.FormValue("font")
	log.Println(text)
	log.Println(font)

	// pass data to ascii-art
	// передает данные в ascii-art
	art, err := ascii.Ascii(text, font)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, funcs.ErrorBad) {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		} else {
			http.Error(w, "Iternal Server Error", http.StatusInternalServerError)
		}
		return

	}

	result := News{
		Body: art,
	}

	// inserts the data "result" into the template
	// вставляет данные в "result" в шаблон
	err = ts.Execute(w, result)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Interal Server Error", 500)
		return
	}
}
