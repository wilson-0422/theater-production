package models

import "time"

type Script struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null;size:200"`
	Author    string    `gorm:"not null;size:100"`
	Genre     string    `gorm:"size:50"`
	Synopsis  string    `gorm:"type:text"`
	Duration  int       `gorm:"default:0"`
	Status    string    `gorm:"size:20;default:'draft'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
