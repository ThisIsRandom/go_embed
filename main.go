package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/thisisrandom/emdedded-rest/database"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	//s := http.NewServeMux()
	d := database.NewDatabase()

	d.Instance.Create(&database.ReadingType{Title: "test"})

	/* readingHandler := handlers.NewReadingsHandler(d.Instance)

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case "POST":
			readingHandler.POST(w, r)
		case "GET":
			readingHandler.GET(w, r)
		}
	})

	url := fmt.Sprint(os.Getenv("IP_ADDR"), ":", os.Getenv("PORT"))
	err := http.ListenAndServe(url, s)
	if err != nil {
		panic(err)
	}
	*/
	fmt.Println("RUNS")
}
