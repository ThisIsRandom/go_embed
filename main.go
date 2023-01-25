package main

import (
	"html/template"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/thisisrandom/emdedded-rest/database"
	"github.com/thisisrandom/emdedded-rest/handlers"
)

func init() {
	godotenv.Load(".env")
}

type TemplateData struct {
	Readings []database.Reading
}

func main() {
	s := http.NewServeMux()
	d := database.NewDatabase()

	readingHandler := handlers.NewReadingsHandler(d.Instance)
	configHandler := handlers.NewConfigsHandler(d.Instance)

	s.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var readings []database.Reading

		if res := d.Instance.Find(&readings); res.Error != nil {
			w.Write([]byte("Error"))
		}

		tmpl, _ := template.ParseFiles("./templates/index.html")

		tmpl.Execute(w, TemplateData{Readings: readings})
	})

	s.HandleFunc("/readings", func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case "POST":
			readingHandler.POST(w, r)
		case "GET":
			readingHandler.GET(w, r)
		}
	})

	s.HandleFunc("/configs", func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case "POST":
			configHandler.POST(w, r)
		case "GET":
			configHandler.GET(w, r)
		}
	})

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		panic(err)
	}
}
