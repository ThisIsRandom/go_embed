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
