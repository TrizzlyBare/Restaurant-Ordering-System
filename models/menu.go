package models

type Menu struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Availability bool   `gorm:"default:true"`
}
