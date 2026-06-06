package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null;size:50"`
	Password  string    `gorm:"not null;size:255"`
	Role      string    `gorm:"not null;size:20;default:'staff'"`
	CreatedAt time.Time
}
