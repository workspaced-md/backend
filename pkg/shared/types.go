package shared

import (
	"time"
)

type Account struct {
	ID        int       `gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

