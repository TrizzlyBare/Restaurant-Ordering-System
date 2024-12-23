package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID 			uint		`gorm:"primaryKey"`
	Name     	string		`gorm:"size 255;not null"`
	Email    	string		`gorm:"size 255;not null;unique"`
	Password 	string		`gorm:"not null"`
	Role    	string		`gorm:"type:enum('admin', 'customer');default:'customer'"`
	CreatedAt	time.Time	`gorm:"autoCreateTime"`
}


var DB *gorm.DB



func InitDB(db *gorm.DB) {

    DB = db

}
