package models

import "time"

type Theater struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"not null;size:200"`
	Location   string    `gorm:"size:200"`
	Capacity   int       `gorm:"default:0"`
	Contact    string    `gorm:"size:100"`
	Facilities string    `gorm:"type:text"`
	CreatedAt  time.Time
}

type TheaterSchedule struct {
	ID        uint      `gorm:"primaryKey"`
	TheaterID uint      `gorm:"not null;index"`
	ScriptID  uint      `gorm:"not null;index"`
	StartTime time.Time
	EndTime   time.Time
	Status    string    `gorm:"size:20;default:'booked'"`
	CreatedAt time.Time
	Theater   Theater   `gorm:"foreignKey:TheaterID"`
	Script    Script    `gorm:"foreignKey:ScriptID"`
}
