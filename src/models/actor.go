package models

import "time"

type Actor struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null;size:100"`
	Gender    string    `gorm:"size:10"`
	Phone     string    `gorm:"size:20"`
	RoleType  string    `gorm:"size:20;column:role_type"`
	Skills    string    `gorm:"type:text"`
	Status    string    `gorm:"size:20;default:'available'"`
	CreatedAt time.Time
}

type ActorSchedule struct {
	ID           uint      `gorm:"primaryKey"`
	ActorID      uint      `gorm:"not null;index"`
	ScriptID     uint      `gorm:"not null;index"`
	RoleName     string    `gorm:"size:100"`
	ScheduleDate time.Time
	Status       string    `gorm:"size:20;default:'scheduled'"`
	CreatedAt    time.Time
	Actor        Actor     `gorm:"foreignKey:ActorID"`
	Script       Script    `gorm:"foreignKey:ScriptID"`
}
