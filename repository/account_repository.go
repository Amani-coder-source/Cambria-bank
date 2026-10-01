package repository

import (
	"database/sql"

	"cambria-bank/models"
)

type AccountRepository struct {
	DB *sql.DB
}

func (r AccountRepository) CreateAccount(account models.Account) error {

	query := `
		INSERT INTO accounts (
			user_id,
			account_number,
			account_type,
			balance,
			currency
		)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		account.UserID,
		account.AccountNumber,
		account.AccountType,
		account.Balance,
		account.Currency,
	)

	return err
}

func (r AccountRepository) GetAccountByUserID(userID int) (models.Account, error) {

	var account models.Account

	query := `
		SELECT
			id,
			user_id,
			account_number,
			account_type,
			balance,
			currency,
			created_at
		FROM accounts
		WHERE user_id = ?
	`

	err := r.DB.QueryRow(query, userID).Scan(
		&account.ID,
		&account.UserID,
		&account.AccountNumber,
		&account.AccountType,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}
