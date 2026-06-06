package models

import "time"

type Prop struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null;size:200"`
	Category    string    `gorm:"size:20;not null"`
	Description string    `gorm:"type:text"`
	Quantity    int       `gorm:"default:0"`
	Available   int       `gorm:"default:0"`
	Status      string    `gorm:"size:20;default:'available'"`
	CreatedAt   time.Time
}

type PropRequisition struct {
	ID              uint      `gorm:"primaryKey"`
	PropID          uint      `gorm:"not null;index"`
	ActorID         uint      `gorm:"not null;index"`
	Quantity        int       `gorm:"default:1"`
	RequisitionDate time.Time
	ReturnDate      time.Time
	Status          string    `gorm:"size:20;default:'pending'"`
	Notes           string    `gorm:"type:text"`
	CreatedAt       time.Time
	Prop            Prop      `gorm:"foreignKey:PropID"`
	Actor           Actor     `gorm:"foreignKey:ActorID"`
}
