package handlers

import (
	"encoding/json"
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

	d := json.NewDecoder(req.Body)

	defer req.Body.Close()

	if err := d.Decode(cfg); err != nil {
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
	}

	if err := h.db.Where("name = ?", cfg.Name).First(&database.Config{}).Error; err != nil {
		h.db.Create(cfg)
	} else {
		h.db.Where("name = ?", cfg.Name).Save(cfg)
	}

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

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)
	res.Write(b)
}
