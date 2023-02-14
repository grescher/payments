package models

import (
	"time"
)

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	IsAdmin      bool   `json:"is_admin"`
	IsActive     bool   `json:"is_active"`
}

type Account struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"` // IBAN
	HolderID int     `json:"user_id"`
	Balance  float64 `json:"balance"`
	IsActive bool    `json:"is_active"`
}

type Payment struct {
	ID          int64     `json:"id"`
	SenderID    int       `json:"sender_id"`
	RecipientID int       `json:"recipient_id"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
	Status      string    `json:"status"`
}

type LogRecord struct {
	ID          int64     `json:"id"`
	Time        time.Time `json:"time"`
	Category    string    `json:"category"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
}
