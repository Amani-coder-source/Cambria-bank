package handlers

import (
	"encoding/json"
	"net/http"

	"cambria-bank/database"
	"cambria-bank/middleware"
	"cambria-bank/repository"
	"cambria-bank/services"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {

	// Get the user ID from the JWT middleware
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)

	if !ok {
		http.Error(
			w,
			"User ID not found",
			http.StatusUnauthorized,
		)
		return
	}

	// Create account repository
	repo := repository.AccountRepository{
		DB: database.DB,
	}

	// Create account service
	service := services.AccountService{
		Repo: repo,
	}

	// Get the logged-in user's account
	account, err := service.GetAccount(userID)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	// Return account as JSON
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(account)
}
