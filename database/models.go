package database

import (
	"gorm.io/gorm"
)

type Reading struct {
	gorm.Model
	Device        string `gorm:"not null"`
	Data          string
	ReadingTypeID string `gorm:"size:256"`
	ReadingType   ReadingType
}

type ReadingType struct {
	Title string `json:"title" gorm:"not null;primaryKey"`
}

type Config struct {
	MinTemp int    `json:"minTemp"`
	MaxTemp int    `json:"maxTemp"`
	MinHum  int    `json:"minHum"`
	MaxHum  int    `json:"maxHum"`
	MaxHour int    `json:"maxHour"`
	MinHour int    `json:"minHour"`
	Name    string `json:"name" gorm:"not null"`
}
