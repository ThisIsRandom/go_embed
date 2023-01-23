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

func main() {
	s := http.NewServeMux()
	d := database.NewDatabase()

	readingHandler := handlers.NewReadingsHandler(d.Instance)
	configHandler := handlers.NewConfigsHandler(d.Instance)

	s.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./static"))))

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/index.html")

		tmpl.Execute(w, nil)
	})

	s.HandleFunc("/readings", func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case "POST":
			readingHandler.POST(w, r)
		case "GET":
			w.Write([]byte("OK"))
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
