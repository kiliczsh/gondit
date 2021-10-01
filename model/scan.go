package model

import "github.com/jinzhu/gorm"

// Scan struct
type Scan struct {
	gorm.Model
	Url       string `gorm:"not null" json:"url"`
	Response  string `gorm:"not null" json:"response"`
}