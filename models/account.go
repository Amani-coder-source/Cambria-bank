package models

type Account struct {
	ID            int     `json:"id"`
	UserID        int     `json:"user_id"`
	AccountNumber string  `json:"account_number"`
	AccountType   string  `json:"account_type"`
	Balance       float64 `json:"balance"`
	Currency      string  `json:"currency"`
	CreatedAt     string  `json:"created_at"`
}
