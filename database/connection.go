package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=tanakrit password=tanakrit dbname=Restaurant-Ordering-System-second port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Fail to connect to database:", err)
    }
    log.Println("Connected to database")
}