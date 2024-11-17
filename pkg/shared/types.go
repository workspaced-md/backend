package shared

import (
	"time"
)

type Account struct {
	ID int
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}