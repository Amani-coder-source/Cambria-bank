package repository

import (
	"database/sql"

	"cambria-bank/models"
)

type TransactionRepository struct {
	DB *sql.DB
}

// CreateTransaction saves a new transaction
func (r TransactionRepository) CreateTransaction(transaction models.Transaction) error {

	query := `
		INSERT INTO transactions (
			account_id,
			transaction_type,
			amount,
			currency,
			description,
			reference,
			status
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		transaction.AccountID,
		transaction.TransactionType,
		transaction.Amount,
		transaction.Currency,
		transaction.Description,
		transaction.Reference,
		transaction.Status,
	)

	return err
}

// GetTransactionsByAccountID gets all transactions
// belonging to one account.
func (r TransactionRepository) GetTransactionsByAccountID(accountID int) ([]models.Transaction, error) {

	query := `
		SELECT
			id,
			account_id,
			transaction_type,
			amount,
			currency,
			description,
			reference,
			status,
			created_at
		FROM transactions
		WHERE account_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(query, accountID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {

		var transaction models.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.TransactionType,
			&transaction.Amount,
			&transaction.Currency,
			&transaction.Description,
			&transaction.Reference,
			&transaction.Status,
			&transaction.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(
			transactions,
			transaction,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
