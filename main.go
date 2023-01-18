package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/thisisrandom/emdedded-rest/database"
	"github.com/thisisrandom/emdedded-rest/handlers"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	log.Println("TESTER")

	s := http.NewServeMux()
	d := database.NewDatabase()

	fmt.Println(d.Instance.Statement.Vars...)

	readingHandler := handlers.NewReadingsHandler(d.Instance)

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case "POST":
			readingHandler.POST(w, r)
		case "GET":
			w.Write([]byte("OK"))
		}
	})

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		panic(err)
	}
}
