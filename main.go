package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/thisisrandom/emdedded-rest/database"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	s := http.NewServeMux()
	d := database.NewDatabase()

	fmt.Println(d.Instance.Statement.Vars...)

	//readingHandler := handlers.NewReadingsHandler(d.Instance)

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case "POST":
			w.Write([]byte("OK"))
		case "GET":
			w.Write([]byte("OK"))
		}
	})

	//url := fmt.Sprint(os.Getenv("IP_ADDR"), ":", os.Getenv("PORT"))

	err := http.ListenAndServe(":8080", s)
	if err != nil {
		panic(err)
	}

	fmt.Println("RUNS")
}
