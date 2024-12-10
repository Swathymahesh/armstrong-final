package database

import (
    "armstrong-webapp/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "root:Password@123@tcp(127.0.0.1:3306)/armstrong_webapp?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate models
    DB.AutoMigrate(&models.User{}, &models.ArmstrongNumber{})
}
