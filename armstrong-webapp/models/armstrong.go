package models

import "time"

type ArmstrongNumber struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    Number    int       `gorm:"not null"`
    CreatedAt time.Time
}
