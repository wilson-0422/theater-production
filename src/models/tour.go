package models

import "time"

type Tour struct {
	ID        uint      `gorm:"primaryKey"`
	ScriptID  uint      `gorm:"not null;index"`
	City      string    `gorm:"size:100"`
	Venue     string    `gorm:"size:200"`
	StartDate time.Time
	EndDate   time.Time
	Status    string    `gorm:"size:20;default:'planned'"`
	Notes     string    `gorm:"type:text"`
	CreatedAt time.Time
	Script    Script    `gorm:"foreignKey:ScriptID"`
}
