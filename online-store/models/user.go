package models

import (
	"time"
)

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `json:"username"`
    Email     string    `gorm:"unique" json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
}
