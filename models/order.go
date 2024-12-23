package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	MenuItems  string    `gorm:"type:json;not null"`
	TotalPrice float64   `gorm:"not null"`
	Status     string    `gorm:"type:enum('Pending', 'Preparing', 'Delivered');default:'Pending'"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

