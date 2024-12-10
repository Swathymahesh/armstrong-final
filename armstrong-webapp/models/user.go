package models

import "time"



type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key field
	 Email string `gorm:"unique;not null" json:"email"`
	 CreatedAt time.Time
 }
 