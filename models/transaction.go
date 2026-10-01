package models

import "time"

type Transaction struct {
	ID              int       `json:"id"`
	AccountID       int       `json:"account_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Currency        string    `json:"currency"`
	Description     string    `json:"description"`
	Reference       string    `json:"reference"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}
