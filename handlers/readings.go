package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thisisrandom/emdedded-rest/database"
	"gorm.io/gorm"
)

type ReadingsHandler struct {
	db *gorm.DB
}

type ReadingPostBody struct {
	Device      string                 `json:"device"`
	Data        map[string]interface{} `json:"data"`
	ReadingType string                 `json:"readingType"`
}

func NewReadingsHandler(db *gorm.DB) *ReadingsHandler {
	return &ReadingsHandler{
		db,
	}
}

func (h *ReadingsHandler) GET(res http.ResponseWriter, req *http.Request) {
	var readings []database.Reading

	if result := h.db.Find(&readings); result.Error != nil {
		res.WriteHeader(500)
		res.Write([]byte(result.Error.Error()))
	}

	b, err := json.Marshal(readings)

	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)
	res.Write(b)
}

func (h *ReadingsHandler) POST(res http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)

	reading := new(ReadingPostBody)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
	}

	defer req.Body.Close()

	json.Unmarshal(b, &reading)

	fmt.Println(string(b))

	r, _ := json.Marshal(reading.Data)

	fmt.Println(string(r))

	m := database.Reading{
		Device: reading.Device,
		Data:   string(r),
		ReadingType: database.ReadingType{
			Title: reading.ReadingType,
		},
	}

	h.db.Create(&m)

	res.Write([]byte("OK"))
}
