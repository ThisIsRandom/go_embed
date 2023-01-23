package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/thisisrandom/emdedded-rest/database"
	"gorm.io/gorm"
)

type ConfigsHandler struct {
	db *gorm.DB
}

func NewConfigsHandler(db *gorm.DB) *ConfigsHandler {
	return &ConfigsHandler{
		db,
	}
}

func (h *ConfigsHandler) POST(res http.ResponseWriter, req *http.Request) {
	cfg := new(database.Config)

	d, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
	}

	if err := json.Unmarshal(d, &cfg); err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
	}

	h.db.Create(cfg)

	res.WriteHeader(200)
	res.Write([]byte("OK"))
}

func (h *ConfigsHandler) GET(res http.ResponseWriter, req *http.Request) {
	var configs []database.Config

	result := h.db.Find(&configs)

	if result.Error != nil {
		res.WriteHeader(500)
		res.Write([]byte(result.Error.Error()))
	}

	b, err := json.Marshal(configs)

	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte(result.Error.Error()))
	}

	res.WriteHeader(200)
	res.Write(b)
}
