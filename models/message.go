package models

import "time"

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	Key       string `gorm:"unique;not null"`
	Value     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
