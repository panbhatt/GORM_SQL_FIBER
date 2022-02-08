package model

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model        // Adds some metadata fields to the table
	ID         string `gorm:"type:string"` // Explicitly specify the type to be uuid
	Title      string `json:"title"`
	SubTitle   string `json:"sub_title"`
	Text       string `json:"text"`
}

type Employee struct {
	gorm.Model
	FName string `gorm:"type:string", json:"firstName`
	LName string `gorm:"type:string", json:"lastName"`
	Age   uint8  `gorm:"type:integer"`
}
