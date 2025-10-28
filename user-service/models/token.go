package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	JTI       string    `json:"jti"` // unique identifier for JWT
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
