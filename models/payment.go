package models

import "time"

type Payment struct {
	ID          uint      `gorm:"primaryKey"`
	OrderID     uint      `gorm:"not null"`
	PaymentMethod string   `gorm:"size:50"`
	Status       string    `gorm:"type:enum('Pending', 'Completed', 'Failed')"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
