package shared

import (
	"time"
)

// Account represents a user account.
type Account struct {
	Id        int       `gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// Workspace represents a workspace.
type Workspace struct {
	Id          int       `gorm:"primaryKey"`
	OwnerId     int       `json:"owner_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
}
