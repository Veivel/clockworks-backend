package models

import "time"

type User struct {
	Username  string    `gorm:"primaryKey" json:"username"`
	CreatedAt time.Time `gorm:"DATETIME" json:"created_at"`
	Email     string    `gorm:"VARCHAR(256)" json:"email"`
	Password  string    `gorm:"VARCHAR(384)" json:"password"`
}
