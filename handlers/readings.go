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
	fmt.Println("test")
	res.Write([]byte("OK"))
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
