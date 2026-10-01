package handlers

import (
	"encoding/json"
	"net/http"

	"cambria-bank/database"
	"cambria-bank/middleware"
	"cambria-bank/models"
	"cambria-bank/repository"
)

// GetTransactions returns the transaction history
// for the authenticated user's account.
func GetTransactions(w http.ResponseWriter, r *http.Request) {

	// Get user ID from JWT
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	// Find the user's bank account
	accountRepo := repository.AccountRepository{
		DB: database.DB,
	}

	account, err := accountRepo.GetAccountByUserID(userID)

	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	// Find transactions belonging to that account
	transactionRepo := repository.TransactionRepository{
		DB: database.DB,
	}

	transactions, err := transactionRepo.GetTransactionsByAccountID(account.ID)

	if err != nil {
		http.Error(w, "Unable to get transactions", http.StatusInternalServerError)
		return
	}

	// Return [] instead of null when there are no transactions
	if transactions == nil {
		transactions = []models.Transaction{}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(transactions)
}
