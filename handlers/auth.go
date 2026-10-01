package handlers

import (
	"encoding/json"
	"net/http"

	"cambria-bank/database"
	"cambria-bank/models"
	"cambria-bank/repository"
	"cambria-bank/services"
)

// Register a new customer and create their bank account
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userRepo := repository.UserRepository{
		DB: database.DB,
	}

	accountRepo := repository.AccountRepository{
		DB: database.DB,
	}

	service := services.AuthService{
		Repo:        userRepo,
		AccountRepo: accountRepo,
	}

	err = service.Register(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Customer and bank account created successfully",
	})

}

// Login an existing customer
func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userRepo := repository.UserRepository{
		DB: database.DB,
	}

	service := services.AuthService{
		Repo: userRepo,
	}

	token, err := service.Login(user.Email, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}
