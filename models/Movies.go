package models

import (
	"time"

	"gorm.io/gorm"
)

type Movies struct {
	Id        uint   `json:"id" gorm:"autoIncrement"`
	Title     string `json:"title" gorm:"type:varchar(255);not null"`
	Slug      string `json:"slug" gorm:"uniqueIndex;type:varchar(255);;not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Duration  int    `json:"duration" gorm:"type:int(5);not null"`
	Image     string `json:"image" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}